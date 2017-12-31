
# Installation Instructions

The first step is different depending on whether you are using MacOS or Windows.

## Step 1 - Install Google Cloud SDK (MacOS)

### Open a Terminal and enter the following: 

>curl <https://sdk.cloud.google.com> | bash

This will install the Google Cloud SDK on your computer. 

During the install, you will be presented with the following message in the Terminal window: 

>Installation directory (this will create a google-cloud-sdk subdirectory) (/Users/[your User here]):

If so, type in the following to continue with the installation process: 

>exec -l $SHELL

This will ensure that your current Terminal is configured properly to run the following commands.

### Run gcloud init:

>gcloud init

NOTE: You will need to enter your Google account user/password details. 

### Install components required by our project:

>gcloud components install app-engine-go

>gcloud components install cloud-datastore-emulator

These components are required to build and test our application.

## Step 1 - Install Google Cloud SDK (Windows)

### Download and run the following installer:

https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe

_Make sure the option to install Bundled Python is checked._

_Accept the options to Start Google Cloud SDK Shell and run gcloud init._

### Install components required for our project:

gcloud components install app-engine-go
gcloud components install cloud-datastore-emulator

### Set GOPATH environment variable:

Add User environment variable GOPATH with value "C:\\Users\\[your User here]\\go".

## Step 2 - Install GitHub Desktop

### Download the appropriate installer from this URL:

https://desktop.github.com/

### Create an account on GitHub.com

Browse to https://github.com and click on "Sign Up".

### Create a local copy of the project source code
1. Browse to https://github.com/mtraudt412/LunchMenu
2. Click on "Clone or download"
3. Click on "Open in Desktop"

You will be prompted for the folder on your computer to clone the project to. This depends on your computer.

####MacOS:

/Users/[your User here]/go/src

####Windows: 

C:\\Users\\[your User here]\go\src

## Step 3 - Install Visual Studio Code

https://code.visualstudio.com/download

## Step 4 - Install Postman

https://www.getpostman.com/apps
