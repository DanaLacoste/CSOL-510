## How To: Use your computer without VirtualBox for CSOL-510

Working with virtual box can be a bit daunting: it simply doesn't work for many people due to compatibility issues and is notorious for compatibility problems and bugs between versions.

But we have to do our homework!  How can we proceed?

Well, for CSOL-510 this is pretty easy: Let's just use our own PCs and get the work done without virtualbox!

### NOTE:

For this document, I am assuming that you will use docker rather than VMs for the assignments.
This is actually what Seed Labs is moving to: version 2 of the Seed Labs content actually uses
docker completely (aside from the base VM, which we will get to below) and containerization has
become quite popular in most corporations: learning how to do it this way can help your understanding
of modern IT departments, but it _is_ different from what the CSOL-510 docs say to do.

## Step 1 : Base OS

### Linux

OK, if you are running Linux for your desktop, you are almost done: the hard part is done anyways!
All you need to do is install docker: `apt-get install docker` if you are on a Debian based distribution (including Ubuntu)

### Windows Laptop (NOT Windows in a VM under Mac, unless you REALLY want to punish yourself!)

Windows has a REALLY amazing feature called "Windows Subsystem for Linux".  To install it, run a PowerShell window AS ADMINISTRATOR and run this command: `wsl --install` (tested on Windows 11 ARM: works fine)

[Microsoft's WSL Installation Instructions](https://docs.microsoft.com/en-us/windows/wsl/install)

After WSL is installed, you may need to do `wsl --install Ubuntu-20.04` or something similar to get the specific distribution you want to work with.

Once it is installed, you can run `wsl --system` and now you are in a Linux system!  Ta-da!  Windows on Linux (and it is actually pretty fast, too).

You can now use either docker for linux (in wsl) or docker for Windows (which has a rather nice GUI) in this configuration:

- `apt-get install docker` from your WSL prompt
- [Get Docker](https://docs.docker.com/get-docker/) for Windows

#### NOTE

I use a Mac at home and Linux at work.   I did NOT proceed past this point in Windows: YMMV.

### Mac

If you have a Mac (INTEL OR M1/Apple Silicon.  YES, BOTH WORK FINE!) then you don't need Linux for this course.  Other courses, maybe, but for this course, all the tools you need are already here.  For the assignments, you can run most of them simply with a terminal, but for the rest you need a few things installed.

I would highly suggest installing HomeBrew, however, before you start: it is BY FAR the easiest way to install most of the additional tools required: [HomeBrew](https://brew.sh/).  (HomeBrew is basically "install linux software on your mac".  It does NOT run linux, it actually installs compiled-for-mac versions of linux tools.)

NOTE: Some of the tools used in the assignments do really basic things: hex edit, view a graphics file, etc.  You can decide if you want to use those tools or if you want to use Mac tools (which are easier to work with, anyways).

Docker for Mac works great, but can conflict with VirtualBox.  You can install it at [Get Docker](https://docs.docker.com/get-docker/).

## OK, I am ready, now what do I need?

This includes what you need for each week's lessons.  Each week is cumulative: week 3 assumes you have week 2 installed, etc.

### Week 2 : Secret-Key Encryption Lab

This week uses the following:

| Tool    | Linux                   | Mac               | Windows (WSL)         | Used in Task  |
|---------|-------------------------|-------------------|-----------------------|---------------|
| openssl | Already Installed       | Already Installed | Already Installed     | 2             |
| eog     | `apt-get install eog`   | Just use Preview  | Just click on the BMP | 3             |
| head    | Already Installed       | Already Installed | Already Installed     | 3             |
| tail    | Already Installed       | Already Installed | Already Installed     | 3             |
| cat     | Already Installed       | Already Installed | Already Installed     | 3             |
| echo    | Already Installed       | Already Installed | Already Installed     | 4             |
| hexdump | Already Installed       | Already Installed | Already Installed     | 4             |
| xxd     | Already Installed       | Already Installed | Already Installed     | 4             |
| bless   | `apt-get install bless` | SEE NOTE BELOW    | SEE NOTE BELOW        |


#### Bless (Hex Editor on Linux)

##### Mac running GUI (X11) Linux Apps

**SPECIAL NOTE**: On a Mac, there is a command named `bless` which is used to sign disks to make them bootable.  Please do not get these mixed up lest you break your system!

You have two options when it comes to this:
- install X so your mac can run linux-style GUI programs [X Quartz](https://www.xquartz.org/) `brew install xquartz`
- Use a different editor, such as [Hex Fiend](https://github.com/HexFiend/HexFiend) (Available in the Mac App Store)

:whynotboth:

##### Windows WSL running GUI (X11) Linux Apps
In WSL, you _can_ run graphical linux programs, but you need to have some things working first: [Run Linux GUI Apps in WSL](https://docs.microsoft.com/en-us/windows/wsl/tutorials/gui-apps) (NOTE: UNTESTED).  Once you do this, `apt-get install bless` should leave you ready to go!

You can alternatively just use a Windows hex editor: there are MANY available [List of hex editors](https://en.wikipedia.org/wiki/Comparison_of_hex_editors).
