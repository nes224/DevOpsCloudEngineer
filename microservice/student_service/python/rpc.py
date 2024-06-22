import psycopg2
from nameko.rpc import rpc

def connect():
    Dbconnect = psycopg2.connect(host='localhost',
                                user='root',
                                passwd='secret',
                                db='student_service',
                                port=5434)
    return Dbconnect

def insert(firstname, lastname, email):
    DBconnect = connect()
    cur = DBconnect.cursor()
    cur.execute("DROP TABLE IF EXISTS student")
    cur.execute("CREATE TABLE student (id int not null, FirstName varchar(255) not null, LastName varchar(255) not null, Email varchar(255) not null);")
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