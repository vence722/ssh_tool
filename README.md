# SSH TOOL

A tool that helps you to use SSH on multiple machines at one time. 

# Introduction

It's a very simple command-line with just an executable file and a configuration file in XML format.

The configuration file consists of two sections --- hosts definition section and commands definition section. Each section provides flexibility for the users to define their behaviors for every single host.

The host definition section defines each host that you want to connect with via SSH. A host is defined as the a struct containing its hostname(or IP address), username and password. Currently only password authentication is supported, and keytab file authentication will be added in the next version.

The commands definition section defines the exact behaviors you want to do with each host. It consists of many sub-sections, each of which represents a command collection corresponding to one or more hosts. The value of the Hostname tag here could be a single hostname, or multiple hostname separated by comma, or an * representing all hosts. Under the Hostname definition, multiple Line tags will be used to define the shell commands that you want to call on the hosts defined above.

# Example

Configuration file example:

	<?xml version="1.0" encoding="UTF-8"?>
	<SSHConfig>
      
	  <!-- hosts definition section -->
	  <Hosts>
	    <Host>
	      <Hostname>H1</Hostname>
	      <User>USERNAME_FOR_H1</User>
	      <Password>PASSWORD_FOR_H1</Password>
	    </Host>
		<Host>
	      <Hostname>H2</Hostname>
	      <User>USERNAME_FOR_H2</User>
	      <Password>PASSWORD_FOR_H2</Password>
	    </Host>
		<Host>
	      <Hostname>H3</Hostname>
	      <User>USERNAME_FOR_H3</User>
	      <Password>PASSWORD_FOR_H3</Password>
	    </Host>
		<Host>
	      <Hostname>H4</Hostname>
	      <User>USERNAME_FOR_H4</User>
	      <Password>PASSWORD_FOR_H4</Password>
	    </Host>
	  </Hosts>
	
	  <!-- commands definition section -->
	  <Commands>
	    <!-- commands for single host -->
	    <Command>
	      <Hostname>H1</Hostname>
	      <Line>cd /etc/</Line>
	      <Line>ls -l</Line>
	    </Command>
		<!-- commands for multiple hosts -->
		<Command>
		  <Hostname>H2,H3</Hostname>
		  <Line>cd /tmp/</Line>
		  <Line>mkdir test</Line>
		  <Line>ls -l</Line>
		</Command>
		<!-- commands for all hosts -->
		<Command>
		  <Hostname>*</Hostname>
		  <Line>echo "Access to the host via SSH!"</Line>
		</Command>
	  </Commands>
	
	</SSHConfig> 
	
Place the configuration file next to your executable, then simply run ./ssh_tool(Unix, Linux, OS X) or ssh_tool.exe(Windows)

## Contributors

[vence722](https://github.com/vence722) (vence722@gmail.com)
