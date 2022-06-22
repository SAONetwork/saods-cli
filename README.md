# saods-cli

The _SAO Data Store Command Line Interface (CLI)_ is a golang application for terminal-based [SAO Data Store] workflows.

Follow the [instruction](https://github.com/SaoNetwork/Data-Store-Guide/blob/main/README.md#authenticate-requests) to get your first application and use CLI to store your first file on the platform with a series of simple CLI commands.

# Getting Started

CLI users and developers must both follow these steps to get the application and running:

## Download bin
Currently, CLI is supported on the following platforms, please choose your bin file:

[windows](https://github.com/SaoNetwork/sao-cli/releases/download/v1.0.0/saods.exe)

[linux](https://github.com/SaoNetwork/sao-cli/releases/download/v1.0.0/saods-linux)

[darwin](https://github.com/SaoNetwork/sao-cli/releases/download/v1.0.0/saods-darwin)


## Set environment
you can set appId and apiKey in environment

## Using the CLI

Learn to use any command:

```shell
saods --help
saods addFile --help
saods getFile --help
saods listFiles --help
```

Here is the example to add file and get file:
```shell
./saods addFile --localPath /path/to/my/file --appId myAppId --apiKey myApiKey
{
    "code":"200",
    "data":{
        "ID":1,
        "CreatedAt":"2022-06-21T14:29:31.689Z",
        "UpdatedAt":"2022-06-21T14:29:31.689Z",
        "appId":"myAppId",
        "filename":"file",
        "contentType":"application/octet-stream",
        "Size":317,
        "ipfsHash":"QmNwB9iCjTwvAS3WjRe1EJzt5E4f966TQGPaNXdCPVCd2b",
        "cid":"",
        "storageProvider":""
    },
    "message":"ok",
    "timestamp":1655821771714
}

./saods getFile --hash myHash --appId myAppId --apiKey myApiKey
--respond with file content
```
