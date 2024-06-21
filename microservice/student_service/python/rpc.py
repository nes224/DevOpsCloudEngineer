import MySQLdb
from nameko.rpc import rpc

def connect():
    Dbconnect = MySQLdb.connect(host='db',
                                user='devops',
                                passwd='devops101',
                                db='devops_db',
                                port=3306)
    return Dbconnect

def insert(firstname, lastname, email):
    DBconnect = connect()
    cur = DBconnect.cursor()
    cur.execute("INSERT INTO student (FirstName, LastName, Email) VALUES (%s,%s,%s);", (firstname,lastname,email))
    id = cur.lastrowid
    DBconnect.commit()
    DBconnect.close()
    return id

class Service:
    name = "student"

    @rpc
    def insert(self, firstname,lastname,email):
        result = insert(firstname,lastname,email)
        return result