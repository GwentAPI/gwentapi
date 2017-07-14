# GwentAPI
How to install and seed the database


## Setup

1. From the project directory, initialize the submodules:
    * ``git submodule init``
    * ``git submodule update``

### Feed the database

The ``manipulator`` submodule contains the software needed to feed the database.

Assuming that you have a running mongoDB instance and you are located in the submodule folder:

1. Compile the program
2. Obtain a copy of the card date in json
3. Run manipulator by inputing the json file ``manipulator -db gwentapi -input data.json``