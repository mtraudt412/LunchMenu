
# Installation Instructions

In the instructions below, any reference to "[your User here]" should be replaced with your Windows or MacOS username.

Follow the instructions in Step 1a or 1b below depending on whether you are using MacOS or Windows. Then follow the instructions for the following steps.

## Step 1a - Install Google Cloud SDK (MacOS)

### Open a Terminal and enter the following: 

>curl <https://sdk.cloud.google.com> | bash

This will install the Google Cloud SDK on your computer. 

During the install, you will be presented with the following message in the Terminal window: 

>Installation directory (this will create a google-cloud-sdk subdirectory) (/Users/[your User here]):

Press the Enter key. This will install the google-cloud-sdk software into the location displayed at the end of the message.

After this step completes, type in the following to continue with the installation process: 

>exec -l $SHELL

This will ensure that your current Terminal is configured properly to run the remaining commands.

### Run gcloud init:

>gcloud init

You will need to enter your Google account user/password details. 

You will be asked if you want to create a new Google App Engine project. You should skip this step.

### Install components required by our project:

>gcloud components install app-engine-go

>gcloud components install cloud-datastore-emulator

These components are required to build and test our application.

## Step 1b - Install Google Cloud SDK (Windows)

### Download and run the following installer:

https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe

_Make sure the option to install Bundled Python is checked._

_Accept the options to Start Google Cloud SDK Shell and run gcloud init._

You will be asked if you want to create a new Google App Engine project. You should skip this step.

### Install components required for our project:

>gcloud components install app-engine-go

>gcloud components install cloud-datastore-emulator

## Step 2 - Install the Go programming language

### Download and run the appropriate installer from this URL:

https://golang.org/dl/

### Set GOPATH environment variable:

Add User environment variable GOPATH with value "C:\\Users\\[your User here]\\go".

* MacOS:

Edit the file /Users/[your User here]/.bash_profile using a Text Editor and add the following line at the bottom:

export GOPATH=/Users/[your User here]/go

* Windows: 

Open the Start Menu and type "Edit Environment Variables For Your Account" and select the menu item that is displayed.

In the upper section of the form that is displayed, click on "New..." and define a new variable with Name="GOPATH" and Value="C:\\Users\\[your User here]\\go".

In the upper section of the form that is displayed, select the row containing the "Path" environment variable, and then click on "New...". This will display a child form. Click on the "New" button on this form and add the following value:

C:\Go\bin

After saving this change, you will need to open a new Command Prompt.

## Step 3 - Install GitHub Desktop

### Download and run the appropriate installer from this URL:

https://desktop.github.com/

### Create an account on GitHub.com

Browse to https://github.com and click on "Sign Up".

### Create a local copy of the project source code
1. Browse to https://github.com/mtraudt412/LunchMenu
2. Click on "Clone or download"
3. Click on "Open in Desktop"

You will be prompted for the folder on your computer to clone the project to. This depends on your computer.

* MacOS:

/Users/[your User here]/go/src

* Windows: 

C:\\Users\\[your User here]\\go\\src

### 4. Install Git command-line (required by Go)

Browse to https://git-scm.com/downloads and install the correct version (MacOS or Windows).

### Test that you are able to build and run the LunchMenu application

Type the following commands into a Terminal:

* MacOS:

cd $GOPATH/src/LunchMenu
./start.sh

* Windows: 

cd %GOPATH%\\src\\LunchMenu
.\start.cmd

At this point, you should open a web browser such as Chrome or Safari and browse to: 

http://localhost:8686

If the application started successfully, you will see the landing page for the LunchMenu application.

## Step 4 - Install Visual Studio Code

https://code.visualstudio.com/download

## Step 5 - Install Postman

https://www.getpostman.com/apps
