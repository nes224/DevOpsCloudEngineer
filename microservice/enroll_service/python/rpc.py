import psycopg2
from nameko.rpc import rpc

def connect():
    DBconnect = psycopg2.connect(host='localhost',
                                user='root',
                                passwd='secret',
                                db='student_service',
                                port=5434)
    
    return DBconnect

def insert(id, firstname, lastname, subjectid, term, year):
    DBconnect = connect()
    cur = DBconnect.cursor()
    cur.execute("DROP TABLE IF EXISTS enroll")
    cur.execute("CREATE TABLE enroll (id int not null,name varchar(255) not null,subjectid varchar(255) not null,term int not null,year int not null);")

    cur.execute("INSERT INTO enroll (id, name, subjectid, term, year) VALUES (%s,%s,%s,%s,%s);", (id,firstname + ''+ lastname, subjectid,term,year))
    id = cur.lastrowid
    DBconnect.commit()
    DBconnect.close()
    return id

class Service:
    name = "enroll"

    @rpc
    def insert(self, id, firstname, lastname):
        result = insert(id, firstname, lastname, '081102', 1, 2563)
        result = insert(id, firstname, lastname, '520101', 1, 2563)
        result = insert(id, firstname, lastname, '511100', 1, 2563)
        result = insert(id, firstname, lastname, '517121', 1, 2563)
        return result