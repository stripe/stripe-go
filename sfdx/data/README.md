# Data related to the Web Service endpoints

* [example-return.json](example-return.json) details the response from  
  the `batchApi` endpoint with a single `order_id` requested.

* [open_api.yaml](open_api.yaml) defines the basic contract of the 
  `batchApi` endpoint, and provides an (known incomplete) view of 
  the shape of the data returned. The objects documented in the YAML 
  file currently are only the ones which have received additional 
  annotated properties. None of the standard fields or unannotated 
  objects have been added to this document yet, as those specifics 
  can be found in the [example-return.json](example-return.json), 
  if needed. 