
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package filestore

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/rand"
	"testing"

	dag "gx/ipfs/QmTQdH4848iTVCJmKXYyRiK72HufWTLYQQ8iN3JaQ8K1Hq/go-merkledag"

	posinfo "gx/ipfs/QmR6YMs8EkXQLXNwQKxLnQp2VBZSepoEJ8KCZAyanJHhJu/go-ipfs-posinfo"
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	blockstore "gx/ipfs/QmS2aqUZLJp8kF1ihE5rvDGE5LvmKDPnx32w9Z1BW9xLV5/go-ipfs-blockstore"
	ds "gx/ipfs/Qmf4xQhNomPNhrtZc67qSnfJSjxjXs9LWvknJtSXwimPrM/go-datastore"
)

func newTestFilestore(t *testing.T) (string, *Filestore) {
	mds := ds.NewMapDatastore()

	testdir, err := ioutil.TempDir("", "filestore-test")
	if err != nil {
		t.Fatal(err)
	}
	fm := NewFileManager(mds, testdir)
	fm.AllowFiles = true

	bs := blockstore.NewBlockstore(mds)
	fstore := NewFilestore(bs, fm)
	return testdir, fstore
}

func makeFile(dir string, data []byte) (string, error) {
	f, err := ioutil.TempFile(dir, "file")
	if err != nil {
		return "", err
	}

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func TestBasicFilestore(t *testing.T) {
	dir, fs := newTestFilestore(t)

	buf := make([]byte, 1000)
	rand.Read(buf)

	fname, err := makeFile(dir, buf)
	if err != nil {
		t.Fatal(err)
	}

	var cids []cid.Cid
	for i := 0; i < 100; i++ {
		n := &posinfo.FilestoreNode{
			PosInfo: &posinfo.PosInfo{
				FullPath: fname,
				Offset:   uint64(i * 10),
			},
			Node: dag.NewRawNode(buf[i*10 : (i+1)*10]),
		}

		err := fs.Put(n)
		if err != nil {
			t.Fatal(err)
		}
		cids = append(cids, n.Node.Cid())
	}

	for i, c := range cids {
		blk, err := fs.Get(c)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(blk.RawData(), buf[i*10:(i+1)*10]) {
			t.Fatal("data didnt match on the way out")
		}
	}

	kch, err := fs.AllKeysChan(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	out := make(map[string]struct{})
	for c := range kch {
		out[c.KeyString()] = struct{}{}
	}

	if len(out) != len(cids) {
		t.Fatal("mismatch in number of entries")
	}

	for _, c := range cids {
		if _, ok := out[c.KeyString()]; !ok {
			t.Fatal("missing cid: ", c)
		}
	}
}

func randomFileAdd(t *testing.T, fs *Filestore, dir string, size int) (string, []cid.Cid) {
	buf := make([]byte, size)
	rand.Read(buf)

	fname, err := makeFile(dir, buf)
	if err != nil {
		t.Fatal(err)
	}

	var out []cid.Cid
	for i := 0; i < size/10; i++ {
		n := &posinfo.FilestoreNode{
			PosInfo: &posinfo.PosInfo{
				FullPath: fname,
				Offset:   uint64(i * 10),
			},
			Node: dag.NewRawNode(buf[i*10 : (i+1)*10]),
		}
		err := fs.Put(n)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, n.Cid())
	}

	return fname, out
}

func TestDeletes(t *testing.T) {
	dir, fs := newTestFilestore(t)
	_, cids := randomFileAdd(t, fs, dir, 100)
	todelete := cids[:4]
	for _, c := range todelete {
		err := fs.DeleteBlock(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	deleted := make(map[string]bool)
	for _, c := range todelete {
		_, err := fs.Get(c)
		if err != blockstore.ErrNotFound {
			t.Fatal("expected blockstore not found error")
		}
		deleted[c.KeyString()] = true
	}

	keys, err := fs.AllKeysChan(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for c := range keys {
		if deleted[c.KeyString()] {
			t.Fatal("shouldnt have reference to this key anymore")
		}
	}
}

func TestIsURL(t *testing.T) {
if !IsURL("http://www.example.com”）；
t.Fatal("IsURL failed: http://www.example.com”）。
	}
if !IsURL("https://www.example.com”）；
t.Fatal("IsURL failed: https://www.example.com”）。
	}
	if IsURL("adir/afile") || IsURL("http:/ /afile") || IsURL("http:/a/file") {
		t.Fatal("IsURL recognized non-url")
	}
}
