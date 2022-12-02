# Test project for company DeNet

## Task
Create CLI application using the go-ethereum library.
Need to implement: 
1) Creating a user and crypto Wallet
2) The password hash must be stored encrypted in the file
3) The data in the hash file must be encrypted using AES

##GUID

###Makefile

USE:

1) make build - to build the project
2) make rebuild - to rebuild the project
3) make clear - to delete a binary file
4) make run - to run the project

At the first launch, the program offers to create an Account. 
To create an account, you only need a password. 
The password must be at least 8 characters long.
The next time you launch the programme use the address you received when creating your account as your login.
