# GwentAPI
How to install and seed the database


## Requirement & dependencies

* Bash (use MINGW/git bash or cygwin for windows)
* MongoDB 3.4
* Python 3.4.2
    * pymongo 3.4.0
* Git
* [jq-1.5](https://stedolan.github.io/jq/)

## Setup

Assuming you installed and configure all the required softwares listed above and that you have a working Go setup:

### Get the source of the software
1. ``go get github.com/GwentAPI/gwentapi``
2. From the project directory, initialize the submodules:
    * ``git submodule init``
    * ``git submodule update``

### Feed the database

The ``data`` submodule contains everything needed to feed the database.

Assuming that you have a running mongoDB instance and you are located in the ``data`` folder:

1. Run the ``extract.sh`` bash script, giving it the latest.jsonl file as a parameter:
    * ``./extract.sh input/latest.jsonl`` which will use jq to extract and create collection related .jsonl files on the same directory as the script.
2. Assuming that ``db_setup.py`` is in the same location as the newly generated files, populate the database by running db_setup.py (the command differs depending on your setup):
    * ``python3 db_setup.py``
    
``db_setup.py`` will create the database if not present along with all the collections. If a collection is already present, the script will refuse to run unless you force it:
* ``python3 db_setup.py -f`` forcing the script will drop all the collections before re-populating the database.

If you enabled authentication on mongoDB or if you are using a different port/host other than the default port on localhost, the script support overriding the values. Learn more with:
``python3 db_setup.py --help``.

SSL is not supported with the script. Remember that the go program **doesn't** support clusters/multiple hosts or SSL. Modify the values inside ``config.toml`` to suit your environment.
