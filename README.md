# SGW - Storage Gateway


Welcome to the Storage Gateway project in Golang! This repository contains the source code and documentation for building a versatile and efficient object storage gateway written in Golang. This gateway serves as a bridge between applications and various object storage providers, allowing seamless data storage and retrieval.

## Features
- **Custom Object Storage URLs**: The gateway can be used to create custom URLs for object storage providers, allowing for more control over the storage and retrieval of data.
- **Versatile Integration**: Connects with popular storage providers, enabling smooth data exchange with cloud-based storage services.
- **Scalability**: Designed with scalability in mind, supporting the growth of data volumes without compromising performance.
- **Secure Communication**: Ensures secure communication through encrypted connections between the gateway and object storage providers.
- **Customizable Configuration**: Offers configuration options to tailor the gateway behavior to your specific use cases.

## Roadmap
- [x] Support for MinIO
- [x] Support for Scaleway Object Storage
- [ ] Support for Storage providers load balancing
- [ ] Support for AWS S3
- [ ] Support for DigitalOcean Spaces
- [ ] Support for Google Cloud Storage
- [ ] Support for Microsoft Azure Blob Storage
- [ ] Support for Alibaba Cloud Object Storage Service
- [ ] Support for IBM Cloud Object Storage
- [ ] Support for Oracle Cloud Infrastructure Object Storage



## Getting Started
These instructions will guide you through setting up and running the Object Storage Gateway on your local machine or within your infrastructure.

## Prerequisites
- Golang v1.21 or higher
- Make

## Installation

### Clone this repository:

```bash
git clone https://github.com/d8x/sgw.git
```

### Navigate to the project directory:

```bash
cd sgw
```

### Build the gateway:

```bash
make build
```

### Adjust the configuration:

```bash
cp sgw.example.yaml sgw.yaml
```

### Run the gateway:

```bash
./sgw
```
Access the gateway at http://localhost:7312 by default.

## Configuration
Rename the sgw.example.yaml file to sgw.yaml.
Open sgw.yaml and configure the storage provider settings.


## Usage
The gateway can be used to create custom URLs for object storage providers, allowing for more control over the storage and retrieval of data.
There are two ways to use the gateway:
- **Using URL Query Parameter**: Append the URL query parameter `provider` to the object storage URL. The value of the parameter should be the name of the storage provider. For example, to use the gateway with MinIO, append `?provider=minioProvider` to the object storage URL.
- **Using Custom Header**: Add a custom header `X-Provider` to the request. The value of the header should be the name of the storage provider. For example, to use the gateway with MinIO, add `X-Storage-Gateway-Provider: minioProvider` to the request header.


## Contributing
We welcome contributions from the community! To contribute to the Storage Gateway project, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Commit your changes with meaningful messages.
4. Push your changes to your fork.
5. Submit a pull request to the main repository.

## License
This project is licensed under the Apache 2.0 License - see the LICENSE file for details.

## Contact
For questions, suggestions, or support, please contact me at d8x[at]d8x.me