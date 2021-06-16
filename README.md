# ydb

ydb (Yadiiiig's Database - If you have a better name please contact me) is a hobby project where I'm trying to make my own NoSQL database. It's still a work in progress, but any feedback is always welcome.

## Main Idea

This is the current idea of how the database will look and be used when everything is finished.
You'll just have to install the executeable or build it yourself. Once you've did this you'll be able to setup an example project using a specific command.
This will setup a folder in a specified path. The folder will contain a data file, structure folder including the table structure and a backup folder.
Within the structure folder you'll be able to setup all the tables you'll need (column name and datatype). (This could also be done using the locally hosted dashboard in the future)
If that is done, you can just run the executable and point it to the database folder and it will run on the designated port.
Afterwards you can import the drivers for the language you are using and use the database.

## Features

The database currently supports basic functionality:

* Select
* Insert
* Update
* Delete

## Upcoming features

* Allowing integers (This shouldn't take that long to implement)
* Adding logging
Starting on making a few drivers:
* Go
* Python
* Java
* Lua
* C#/.NET
Creating a local dashboard to:
* View, edit, delete data manually
* Create new database
* Create new tables
* View statistic
* ...

### People who helped me a lot

* [sn0w](https://memleak.eu/sn0w)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
