package raftkv

import (
	"crypto/rand"
	"labrpc"
	"math/big"
	"time"
)

const RetryInterval = time.Duration(125 * time.Millisecond)

type Clerk struct {
	servers []*labrpc.ClientEnd
	// You will have to modify this struct.
	clientId   int64
	RequestSeq int
	leaderId   int
}

func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	x := bigx.Int64()
	return x
}

func MakeClerk(servers []*labrpc.ClientEnd) *Clerk {
	ck := new(Clerk)
	ck.servers = servers
	// You'll have to add code here.
	ck.clientId = nrand()
	ck.RequestSeq = 0
	ck.leaderId = 0
	return ck
}

func (ck *Clerk) Call(rpcname string, args interface{}, reply interface{}) bool {
	return ck.servers[ck.leaderId].Call(rpcname, args, reply)
}

//
// fetch the current value for a key.
// returns "" if the key does not exist.
// keeps trying forever in the face of all other errors.
//
// you can send an RPC with code like this:
// ok := ck.servers[i].Call("KVServer.Get", &args, &reply)
//
// the types of args and reply (including whether they are pointers)
// must match the declared types of the RPC handler function's
// arguments. and reply must be passed as a pointer.
//
func (ck *Clerk) Get(key string) string {

	// You will have to modify this function.
	args := GetArgs{key}
	for {
		var reply GetReply
		if ck.Call("KVServer.Get", &args, &reply) && reply.Err == OK {
			DPrintf("[%d GET key %s reply %#v]", ck.leaderId, args.Key, reply)
			return reply.Value
		}
		ck.leaderId = (ck.leaderId + 1) % len(ck.servers)
		time.Sleep(RetryInterval)
	}
	return ""
}

//
// shared by Put and Append.
//
// you can send an RPC with code like this:
// ok := ck.servers[i].Call("KVServer.PutAppend", &args, &reply)
//
// the types of args and reply (including whether they are pointers)
// must match the declared types of the RPC handler function's
// arguments. and reply must be passed as a pointer.
//
func (ck *Clerk) PutAppend(key string, value string, op string) {
	// You will have to modify this function.
	ck.RequestSeq++
	args := PutAppendArgs{key, value, op, ck.clientId, ck.RequestSeq}
	for {
		var reply PutAppendReply
		if ck.Call("KVServer.PutAppend", &args, &reply) && reply.Err == OK {
			if op == "Put" {
				DPrintf("[%d seq %d PUT key %s value %s reply %v]", ck.leaderId, args.RequestSeq, args.Key, args.Value, reply)
			} else {
				DPrintf("[%d seq %d APPEND key %s value %s reply %v]", ck.leaderId, args.RequestSeq, args.Key, args.Value, reply)
			}
			return
		}
		ck.leaderId = (ck.leaderId + 1) % len(ck.servers)
		time.Sleep(RetryInterval)
	}
}

func (ck *Clerk) Put(key string, value string) {
	ck.PutAppend(key, value, "Put")
}
func (ck *Clerk) Append(key string, value string) {
	ck.PutAppend(key, value, "Append")
}
