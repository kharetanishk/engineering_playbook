Package 
A package is a bundle containing:

The software itself (executables)
Libraries it needs
Configuration files
Metadata (version, dependencies, description)
Installation and removal instructions

For example, Git is distributed as a package.

Package manager

A package manager is a program that automates:

Downloading software
Installing it
Updating it
Removing it
Resolving dependencies
Knows where to put all files in the linux system
Ensure the integrity and authenticity of the package 

Instead of downloading installers from websites like on Windows, Linux uses package managers.

in linux - we have APT ( advance package Tool) package manager ,apt-get(donot give more user friendly output) , snap(self contained dependencies , automatic updates , larger files , universal linux distribution support , does not use storage efficiently  ) | apt(dependencies are shared across files , manual updates , smaller files , only specifc linux distributions  , uses storage efficiently)

APT usage

sudo apt search [package_name] - serach the packages 
or just write java , go , node - it will automatically say 
that the command not found just install with these [suggestions]

to uninstall - sudo apt remove package_name

package manager fetches the softwares from repositories

sudo apt update - update the package index -> pull the latest changes from the apt repo so that we get newer version of repo


2 major linux distributions 

debian based - ubuntu , debian , mint -| package manager -> apt, apt-get
redhat based - RHEL , CentoS , Fedora -| package manager -> yum