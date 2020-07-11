import grpc
import json
import logging
import service_pb2
import service_pb2_grpc

def setRecordResult(stub,SetPerson):
    result = stub.RecordDB(SetPerson)
    if not result.success:
        print("Kullanıcı Kaydedilemedi.")
        return
    if result.success:
        print("Kullanıcı kaydedildi Key: %s Value: %s" % (SetPerson.key, SetPerson.value))
    
def recordResult(stub,data,i):
    setRecordResult(stub, service_pb2.SetPerson(key = i, value = data))

def getGetRecordResult(stub,GetPerson):
    result = stub.GetRecordDB(GetPerson)
    if not result.success:
        print("Kullanıcı Bulunamadı.")
        return
    if result.success:
        print("Kullanıcı Bulundu key: %s " % (GetPerson.key))

def getRecordResult(stub,i):
    getGetRecordResult(stub, service_pb2.GetPerson(key = i))

def readJSON(jsonDosyasi):
    with open(jsonDosyasi) as f:
        data = json.load(f)
        return data
def run():
    with grpc.insecure_channel('localhost:4040') as channel:
        stub = service_pb2_grpc.GrpcServiceStub(channel)
        for i in range(1, 11):
            jsonDosyasi = str(i) + ".json"
            data = readJSON("users/" + jsonDosyasi)
            print("---------------Record Result---------------")    
            recordResult(stub,str(data),str(i))
        for i in range(1, 11):
            print("---------------Get Record Result---------------")
            getRecordResult(stub,str(i))

if __name__ == '__main__':
    logging.basicConfig()
    run()