# APIC Profiler
This tool was created to extract non runtime statistics from an API Connect 2018 system and as a sample for how to call API Connects provider APIs.

To run the tool either go to the releases page or to build from source. To build from source clone the repo and run make in the root directory.

A Sample default output is below this information.

## Usage
Please note if parameters are not set the user is prompted for the required values.

```
Chriss-MacBook-Pro-2:APIC-Profiler chris$ ./api-profiler-osx-v1.01-1561420129  --help
Usage of ./out/v1.01-1561420129/api-profiler-osx-v1.01-1561420129:
  -APIMorg string
    	Organiztion List to investigate. Please multiple orgs in csv, e.g. dev,test,chrisp,marketting
  -APIMpassword string
    	APIConnect Password for utils.Logging into the API Manager Endpoint, if this is not set it is prompted
  -APIMrealm string
    	Realm for utils.Logging into the API Manager Endpoint, if this is not set it is prompted
  -APIMuser string
    	APIConnect User for utils.Logging into the API Manager Endpoint, if this is not set it is prompted
  -CMpassword string
    	APIConnect Password for utils.Logging into the Cloud Manager Endpoint, if this is not set it is prompted
  -CMrealm string
    	Realm for utils.Logging into the Cloud Manager Endpoint, if this is not set it is prompted
  -CMuser string
    	APIConnect User for utils.Logging into the Cloud Manager Endpoint, if this is not set it is prompted
  -debug
    	Enable trace for this appication
  -output string
    	Dictates the output type for the script. Value must be one of é table ö json ö yaml ö verbose ê  (default "table")
  -server string
    	APIConnect Cloud endpoint, if this is not set it is prompted (default "unset")
```


## Sample default Output
--------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------

# Cloud Name = 'Cloud Admin'
-1 means an issue retreiving the data


--------------------------------------------------------------------------------------------------------

 Cloud Manager Details

	 DONE chrisp2
|      NUMBER OF ORGS       | 17 |
|---------------------------|----|
| Number of Members         |  2 |
| Number of Member Invites  |  0 |
| Number of Oauth Providers |  0 |
| Number of User Registries |  3 |
| Number of Mail Servers    |  1 |

--------------------------------------------------------------------------------------------------------

 ## Number of Availability Zones 2
--------------------------------------------------------------------------------------------------------
|           NAME            | NUMBER OF V5C GWS | NUMBER OF API GWS | NUMBER OF PORTALS | NUMBER OF ANALYTICS |
|---------------------------|-------------------|-------------------|-------------------|---------------------|
| availability-zone-default |                 1 |                 1 |                 1 |                   1 |
| performance-testing       |                 1 |                 1 |                 0 |                   0 |
|                           |                   |                   |                   |                     |

--------------------------------------------------------------------------------------------------------

## Number of Provider Organizations to investigate 4
--------------------------------------------------------------------------------------------------------
|    NAME     | MEMBERS | MEMBERINVITES | DRAFTAPIS | AVG API SIZE | MAX API SIZE | DRAFTPRODUCTS | TLS PROFILES | OAUTH PROIVDERS | USER REGISTRIES | KEYSTORES | TRUSTSTORES | CATALOGS |
|-------------|---------|---------------|-----------|--------------|--------------|---------------|--------------|-----------------|-----------------|-----------|-------------|----------|
| Performance |       3 |             0 |         4 |         3143 |         3402 |             5 |            3 |               0 |               2 |         0 |           0 |        1 |
| chrisp      |       2 |             0 |         5 |        43589 |        91159 |            28 |            3 |               2 |               4 |         0 |           0 |        3 |
| chrisp3     |       1 |             0 |         8 |       146820 |       941383 |             7 |            3 |               2 |               2 |         0 |           0 |        1 |
| chrisp2     |       1 |             0 |         0 |            0 |            0 |             0 |            0 |               0 |               0 |         0 |           0 |        1 |
|             |         |               |           |              |              |               |              |                 |                 |           |             |          |

--------------------------------------------------------------------------------------------------------

## 		 Number of catalogs investigated - 6
--------------------------------------------------------------------------------------------------------
|     ORG     |             CATALOG NAME             | MEMBERS | MEMBERINVITES | APIS | AVG API SIZE | MAX API SIZE | PRODUCTS | CONSUMERORGS | PORTAL | TLSPROFILES | OAUTHPROS | USERREGS | SPACES | APPS | SUBSCRIPTIONS | WEBHOOKS |
|-------------|--------------------------------------|---------|---------------|------|--------------|--------------|----------|--------------|--------|-------------|-----------|----------|--------|------|---------------|----------|
|             |                                      |         |               |      |              |              |          |              |        |             |           |          |        |      |               |          |
| Performance | Sandbox                              |       3 |             0 |    4 |         3948 |         4099 |        5 |            1 | false  |           1 |         0 |        0 |      0 |    1 |             3 |        3 |
| chrisp      | a6f7f15e-0290-40af-afb5-ab7ab998a484 |       2 |             0 |    0 |            0 |            0 |        0 |            0 | false  |           0 |         0 |        0 |      0 |    0 |             0 |        3 |
| chrisp      | Sandbox                              |       2 |             0 |    4 |        31711 |        92334 |        4 |            3 | true   |           1 |         0 |        0 |      0 |    1 |             2 |        3 |
| chrisp      | spacetes asdfa;dfauidh adfafd        |       2 |             0 |    3 |        29641 |        45288 |        1 |            0 | false  |           2 |         0 |        0 |      3 |    0 |             0 |        2 |
| chrisp3     | Sandbox                              |       1 |             0 |    5 |       220769 |       943514 |        6 |            1 | false  |           1 |         1 |        0 |      0 |    2 |             5 |        3 |
| chrisp2     | Sandbox                              |       1 |             0 |    0 |            0 |            0 |        0 |            0 | false  |           0 |         0 |        0 |      0 |    0 |             0 |        0 |

--------------------------------------------------------------------------------------------------------

## 		 Number of Webhooks investigated - 14
--------------------------------------------------------------------------------------------------------
|              WEBHOOK ID              |              ORGANIZATION              |                CATALOG                 |       STATE       | LEVEL |            TITLE            |        CREATED AT        |        UPDATED AT        |
|--------------------------------------|----------------------------------------|----------------------------------------|-------------------|-------|-----------------------------|--------------------------|--------------------------|
|                                      |                                        |                                        |                   |       |                             |                          |                          |
| f1f29d2f-547a-47cb-ad05-56c9c2911fd4 | Performance                            | Sandbox                                | online_configured | sync  | idg77v-v5c                  | 2019-04-26T13:17:01.770Z | 2019-04-26T13:17:15.389Z |
|                                      | (3c653952-d3e2-4284-8c83-2d48e896b4fe) | (1b144fd2-f0fc-49d3-acab-730beb28ebcd) |                   |       |                             |                          |                          |
| c5de4a53-a79f-4433-bbdb-fed3c3db35f8 | Performance                            | Sandbox                                | online_configured | sync  | Performance Testing Gateway | 2019-04-25T19:59:01.899Z | 2019-04-29T16:06:41.719Z |
|                                      | (3c653952-d3e2-4284-8c83-2d48e896b4fe) | (1b144fd2-f0fc-49d3-acab-730beb28ebcd) |                   |       |                             |                          |                          |
| e7c88fc2-e986-4e85-93b8-c455ad07daaa | Performance                            | Sandbox                                | online_configured | sync  | Performance Testing v5c     | 2019-04-26T14:13:22.073Z | 2019-04-29T16:06:41.814Z |
|                                      | (3c653952-d3e2-4284-8c83-2d48e896b4fe) | (1b144fd2-f0fc-49d3-acab-730beb28ebcd) |                   |       |                             |                          |                          |
| 4162ac86-ec78-4237-9aa0-345acc24278d | chrisp                                 | a6f7f15e-0290-40af-afb5-ab7ab998a484   | online_configured | sync  | idg77v                      | 2019-04-26T10:42:58.826Z | 2019-04-26T10:43:15.574Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (a6f7f15e-0290-40af-afb5-ab7ab998a484) |                   |       |                             |                          |                          |
| 96daa6b3-81ed-448e-8c21-1994a08165ad | chrisp                                 | a6f7f15e-0290-40af-afb5-ab7ab998a484   | online_configured | sync  | idg77v-v5c                  | 2019-04-26T10:43:03.193Z | 2019-04-26T10:43:15.650Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (a6f7f15e-0290-40af-afb5-ab7ab998a484) |                   |       |                             |                          |                          |
| d8d14f26-7a09-47db-8998-c0b5313bf944 | chrisp                                 | a6f7f15e-0290-40af-afb5-ab7ab998a484   | online_configured | sync  | Performance Testing Gateway | 2019-04-26T10:43:45.085Z | 2019-04-29T16:06:35.458Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (a6f7f15e-0290-40af-afb5-ab7ab998a484) |                   |       |                             |                          |                          |
| 2bddcbbb-ac02-4d9b-98ce-898030d8b121 | chrisp                                 | Sandbox                                | online            | sync  | IBM Developer Portal        | 2019-04-24T11:12:10.119Z | 2019-05-01T13:06:14.483Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (bf4c0ec9-95c8-4e71-a2fd-fec37bbb79d1) |                   |       |                             |                          |                          |
| 158bd1cd-280f-4939-93ef-8479a7f4a99b | chrisp                                 | Sandbox                                | online            | sync  | idg77v                      | 2019-03-04T14:19:05.289Z | 2019-03-21T16:33:06.720Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (bf4c0ec9-95c8-4e71-a2fd-fec37bbb79d1) |                   |       |                             |                          |                          |
| f510da41-d49e-4c97-9814-3f4beb493baf | chrisp                                 | Sandbox                                | online            | sync  | idg77v-v5c                  | 2019-03-08T17:38:45.159Z | 2019-03-21T18:09:43.892Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (bf4c0ec9-95c8-4e71-a2fd-fec37bbb79d1) |                   |       |                             |                          |                          |
| 082ba1cd-c41c-4684-9e96-0adfb6bfda61 | chrisp                                 | spacetes asdfa;dfauidh adfafd          | online_configured | sync  | idg77v                      | 2019-04-24T11:55:42.163Z | 2019-04-24T11:56:00.152Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (42298b69-0edf-4cc3-83b4-232786976bbe) |                   |       |                             |                          |                          |
| 125a63ff-699c-4249-ba7b-9557eb9e09d5 | chrisp                                 | spacetes asdfa;dfauidh adfafd          | online_configured | sync  | idg77v-v5c                  | 2019-04-24T11:55:47.721Z | 2019-04-24T11:56:00.247Z |
|                                      | (8643dbab-b431-4602-a083-8d8fa29d2f6e) | (42298b69-0edf-4cc3-83b4-232786976bbe) |                   |       |                             |                          |                          |
| adac2816-3091-47d0-81d7-02e16718146e | chrisp3                                | Sandbox                                | online_configured | sync  | idg77v                      | 2019-03-25T08:59:18.597Z | 2019-03-25T08:59:23.015Z |
|                                      | (b7cc4351-6be1-4047-a645-b1cc01942c9a) | (6db92f5f-2f10-4d91-bd20-98427ff58f5c) |                   |       |                             |                          |                          |
| f1d48c28-d420-4d01-ba70-4aea110323b4 | chrisp3                                | Sandbox                                | online_configured | sync  | idg77v-v5c                  | 2019-03-25T08:59:25.592Z | 2019-03-25T08:59:45.223Z |
|                                      | (b7cc4351-6be1-4047-a645-b1cc01942c9a) | (6db92f5f-2f10-4d91-bd20-98427ff58f5c) |                   |       |                             |                          |                          |
| 7c16e811-dc9e-4897-873c-39798cc3a157 | chrisp3                                | Sandbox                                | online_configured | sync  | Performance Testing Gateway | 2019-04-26T10:44:17.148Z | 2019-04-29T16:06:46.215Z |
|                                      | (b7cc4351-6be1-4047-a645-b1cc01942c9a) | (6db92f5f-2f10-4d91-bd20-98427ff58f5c) |                   |       |                             |                          |                          |
