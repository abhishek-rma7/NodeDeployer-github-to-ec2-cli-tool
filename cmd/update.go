package cmd

import (
	"NodeDeployer/config"
	"NodeDeployer/internal/git"
	"NodeDeployer/internal/ssh"
	"NodeDeployer/internal/utils"
	"fmt"

	"github.com/urfave/cli/v2"
)

func Update(c *cli.Context) error {
	cfg, err := config.GetConfig(c)
	if err != nil {
		utils.ErrorLogger.Println("Configuration error:", err)
		return err
	}

	utils.InfoLogger.Println("Validating EC2 instance setup...")
	validateCmd := "command -v git && command -v npm && command -v pm2"
	if err := ssh.RunRemoteCommand(cfg.EC2User, cfg.EC2Address, cfg.EC2KeyPath, validateCmd); err != nil {
		utils.ErrorLogger.Println("Validation failed:", err)
		return fmt.Errorf("validation failed: %w", err)
	}

	utils.InfoLogger.Println("Pushing changes to GitHub...")
	if err := git.PushChanges(); err != nil {
		utils.ErrorLogger.Println("Failed to push changes:", err)
		return fmt.Errorf("failed to push changes: %w", err)
	}

	utils.InfoLogger.Println("Creating rollback point...")
	createRollbackCmd := fmt.Sprintf("cd %s && git tag rollback-point", cfg.RepoPath)
	if err := ssh.RunRemoteCommand(cfg.EC2User, cfg.EC2Address, cfg.EC2KeyPath, createRollbackCmd); err != nil {
		utils.ErrorLogger.Println("Failed to create rollback point:", err)
		return fmt.Errorf("failed to create rollback point: %w", err)
	}

	utils.InfoLogger.Println("Pulling latest changes on EC2 instance...")
	pullCmd := fmt.Sprintf("cd %s && git pull && npm install && npm run build && pm2 restart all", cfg.RepoPath)
	if err := ssh.RunRemoteCommand(cfg.EC2User, cfg.EC2Address, cfg.EC2KeyPath, pullCmd); err != nil {
		utils.ErrorLogger.Println("Failed to update EC2 instance:", err)
		rollback(cfg)
		return fmt.Errorf("failed to update EC2 instance: %w", err)
	}

	utils.InfoLogger.Println("Checking application status...")
	statusCmd := fmt.Sprintf("cd %s && pm2 status", cfg.RepoPath)
	if err := ssh.RunRemoteCommand(cfg.EC2User, cfg.EC2Address, cfg.EC2KeyPath, statusCmd); err != nil {
		utils.ErrorLogger.Println("Failed to check application status:", err)
		rollback(cfg)
		return fmt.Errorf("failed to check application status: %w", err)
	}

	utils.InfoLogger.Println("Deployment completed successfully.")
	return nil
}

func rollback(cfg *config.Config) {
	utils.InfoLogger.Println("Rolling back to the previous version...")
	rollbackCmd := fmt.Sprintf("cd %s && git checkout rollback-point && npm install && npm run build && pm2 restart all", cfg.RepoPath)
	if err := ssh.RunRemoteCommand(cfg.EC2User, cfg.EC2Address, cfg.EC2KeyPath, rollbackCmd); err != nil {
		utils.ErrorLogger.Println("Failed to rollback:", err)
	} else {
		utils.InfoLogger.Println("Rollback completed successfully.")
	}
}
