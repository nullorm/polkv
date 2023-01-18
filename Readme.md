# polkv

## index

index.chain.0 - 64b; 64b; 64b; 64b; 64b; --> 10k entries  
index.chain.1 --> 1k entries  
index.linear-probe.0 --> 10k entries linear probing

## data

have 2 files each of block sizes

- 80b - data.80b.0 data.80b.1
- 800b - data.800b.0 data.800b.1
- 8kb - data.8kb.0 data.8kb.1
- 80kb - data.80kb.0 data.80kb.1
- 800kb - data.800kb.0 1.800kb

defragmentor (in background) will run every x mins on a single file at a time after locking the file for writing.  
ensure that when a defragment job runs on data.80b.0, data.80b.1 is open for writing and vice versa.  
all writes will happen in the alternate file while defragmenter has a lock on the file.
note - figure out if this can be avoided.

data to store

- entry is deleted or not - 1b
- key value - 64b
- data block - 80b to 800kb based on file
