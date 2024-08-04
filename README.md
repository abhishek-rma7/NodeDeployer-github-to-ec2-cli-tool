# NodeDeployer

<p align="center">
  <img src="https://i.ibb.co/PzhMsdt/Screenshot-2024-08-04-at-6-43-58-PM.png" alt="NodeDeployer Logo" width="200" style="object-fit: contain">
</p>

<p align="center">
  <strong>Streamline your Node.js deployments from GitHub to AWS EC2 with ease!</strong>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#future-work">Future Work</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Node.js-339933?style=for-the-badge&logo=nodedotjs&logoColor=white" alt="Node.js">
  <img src="https://img.shields.io/badge/Amazon_AWS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white" alt="AWS">
  <img src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white" alt="GitHub">
</p>

NodeDeployer is a powerful and easy-to-use CLI tool designed to streamline the deployment of Node.js projects from GitHub to AWS EC2 instances. With built-in features for logging, configuration management, and rollback support, NodeDeployer makes it easy to manage your deployments efficiently and reliably.

## ✨ Features

- 🚀 **Easy Configuration**: Set up your deployment with simple command-line flags or a JSON config file
- 📝 **Detailed Logging**: Keep track of every step with comprehensive deployment logs
- ⏪ **Rollback Support**: Automatically revert to a previous state if deployment fails
- 🔍 **Status Check**: Verify your deployed application's status with ease
- ✅ **Validation**: Ensure all necessary tools are available on your EC2 instance before deployment

## 🛠️ Installation

Get started with NodeDeployer in just a few steps:

```sh
git clone https://github.com/abhishek-rma7/NodeDeployer-github-to-ec2-cli-tool.git
cd NodeDeployer
go build -o NodeDeployer
```

# 🚀 Usage
NodeDeployer offers flexibility in how you use it. Here are some examples:
## Using Command Line Flags
```sh
  ./NodeDeployer update --ec2-user ec2-user --ec2-address your-ec2-instance-address --ec2-key-path /path/to/your/key.pem --repo-path /path/to/your/repo
```
## Using a Configuration File
```sh
{
  "ec2_user": "ec2-user",
  "ec2_address": "your-ec2-instance-address",
  "ec2_key_path": "/path/to/your/key.pem",
  "repo_path": "/path/to/your/repo"
}
```
Then run:

``` sh
./NodeDeployer update --config path/to/config.json
```
# ⚙️ Configuration
NodeDeployer requires the following parameters:

- ec2_user: EC2 instance username
- ec2_address: EC2 instance address
- ec2_key_path: Path to the SSH key file
- repo_path: Path to the repository on the EC2 instance

# 📊 Logging
All actions are logged to deploy.log in the current directory, providing detailed information about each deployment step.

# 🔄 Rollback Support
In case of deployment failure, NodeDeployer automatically reverts the EC2 instance to the last known good state, ensuring your application remains available.

# 🚧 Future Work

We're constantly improving NodeDeployer. Here's what's on our roadmap:

- 🌐 Multi-Instance Support
- 🔁 Advanced Rollback Strategies
- 🔗 CI/CD Integration
- 🔒 Enhanced Security Features
- 🌍 Support for Other Programming Languages

# 🤝 Contributing
We welcome contributions! If you have ideas for new features or improvements, please open an issue or submit a pull request.

# 📄 License
NodeDeployer is licensed under the MIT License.

# 📞 Contact

For questions or suggestions, reach out to solank45@uwindsor.ca.
<p align="center">
  Made with ❤️ by Abhishek Solanki
</p>
