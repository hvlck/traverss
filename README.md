# traverss

for when you need some json with your rss

Traverss is an easily-deployable utility to convert RSS or Atom feeds into JSON feeds.

## Usage

The main endpoint is `/json/{url}`.

**Do not include a protocol with the URL.**

`/json/drewdevault.com/blog/index.xml` is acceptable, but `/json/https://drewdevault.com/blog/index.xml` is not!

The resulting response body will be the JSON feed of the provided URL.

Responses look like

```json
{
    "title": "",
    "description": "",
    "link": "",
    "feedLink": "",
    "links": [
        ""
    ],
    "updated": "",
    "updatedParsed": "",
    "generator": "",
    "extensions": {
        "atom": {
            "link": [
                {
                    "name": "",
                    "value": "",
                    "attrs": {
                        "href": "",
                        "rel": "self",
                        "type": "application/rss+xml"
                    },
                    "children": {}
                }
            ]
        }
    },
    "items": [
        {
            "title": "",
            "description": "<content>",
            "link": "",
            "links": [
                ""
            ],
            "published": "",
            "publishedParsed": "",
            "guid": ""
        },
        // ...
    ],
}
```

### Errors

Errors are always in the form of JSON, and look like this:

```json
{
    "status": "failed",
    "error": "<Error>"
}
```

Where `<Error>` is one of the following:

+ `not rss` (see below)
+ `failed to fetch url` (invalid URL/request failed/etc)
+ `failed to parse feed` (feed is invalid or unable to be parsed)

#### 'not rss'

The API will fail if the provided URL does not resolve to a `content-type` header of one of the following:

+ `application/atom+xml`
+ `text/xml`
+ `application/xml`
+ `application/rss+xml`
+ `application/feed+json`
