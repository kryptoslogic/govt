govt - VirusTotal API for Go
============================

`govt` is a go module to use the [API](https://www.virustotal.com/documentation/public-api/) of [VirusTotal.com](https://www.virustotal.com/).


Implemented Features
====================

| *Resource* | *Description* | *VT API* | *govt support* |
|------------------------------------|----------------------------------------------------------------------------------------|-------|-----|
| POST /vtapi/v2/file/scan           | Upload a file for scanning with VirusTotal.                                            | public |true |
| GET /vtapi/v2/file/scan/upload_url | Get a special URL to upload files bigger than 32MB in size.                            | private|false|
| POST /vtapi/v2/file/rescan         | Rescan a previously submitted file or schedule a scan to be performed in the future.   | public |true |
| POST /vtapi/v2/file/rescan/delete  | Delete a previously scheduled scan.                                                    | private|false|
| GET /vtapi/v2/file/report          | Get the scan results for a file.                                                       | public |true |
| GET /vtapi/v2/file/behaviour       | Get a report about the behaviour of the file when executed in a sandboxed environment. | private|true |
| GET /vtapi/v2/file/network-traffic | Get a dump of the network traffic generated by the file when executed.                 | private|true |
| GET /vtapi/v2/file/search          | Search for samples that match certain binary/metadata/detection criteria.              | private|true |
| GET /vtapi/v2/file/clusters        | List file similarity clusters for a given time frame.                                  | private|false|
| GET /vtapi/v2/file/distribution    | Get a live feed with the lastest files submitted to VirusTotal.                        | private|true |
| GET /vtapi/v2/file/download        | Download a file by its hash.                                                           | private|true |
| GET /vtapi/v2/file/false-positives | Consume file false positives from your notifications pipe.                             | private|false|
| POST /vtapi/v2/url/scan            | Submmit a URL for scanning with VirusTotal.                                            | public |true |
| GET /vtapi/v2/url/report           | Get the scan results for a given URL.                                                  | public |true |
| GET /vtapi/v2/url/distribution     | Get a live feed with the lastest URLs submitted to VirusTotal.                         | private|false|
| GET /vtapi/v2/ip-address/report    | Get information about a given IP address.                                              | public |true |
| GET /vtapi/v2/domain/report        | Get information about a given domain.                                                  | public |true |
| POST /vtapi/v2/comments/put        | Post a comment on a file or URL.                                                       | public |true |
| GET /vtapi/v2/comments/get         | Get comments for a file or URL.                                                        | private|true |

Missing Features
================

- all of the above with a `false` in the `govt support` column.
- at least for testing the VT apikey has currently be put into the source (get the apikey from a file or an environment variable would be better)
- more and better testing

Install
=======

If you have a go workplace setup and working you can simply do:

 ```go get github.com/williballenthin/govt```

 ```go install github.com/williballenthin/govt```

Usage
=====

In order how to use the `govt` module please have a look at the `SampleClients` directory and it's content.

You need to have an VirusTotal API Key. You can register for an account at VirusTotal in order to get an public API key.
There are also private API keys available, for those you have to be accepted by VirusTotal and you need to pay for.
Depending on your API Key and the access level granted you can use all of the above functions, all but the ones reserved for AV companies, or just the public ones (if you have a free publich API key).

Check out the [README.md](https://github.com/williballenthin/govt/blob/master/SampleClients/README.md) file in the `SampleClients` directory to find out how to set-up your API key in order to use the provided Example programs.

Authors
=======

`govt` was initially written by `Willi Ballenthin`. Later improved and new features added by `Christopher 'tankbusta' Schmitt` and `Florian 'scusi' Walther`
