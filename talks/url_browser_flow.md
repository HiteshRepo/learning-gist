## URL Components
Given URL: http://example.com/product/electric/phone
1. Scheme: 'http://' -> This tells the browser to connect to server using a protocol called 'http'. Another common scheme is 'https', where the connection is encrypted.
2. Domain: 'example.com' -> Domain name of the site.
3. Path: 'product/electric'.
4. Resource: 'phone'.

The difference between 'Path' and "Resource" is not always clear. But it is analogous to directories and files.
They together specify the resource on the server we would like to load.

## What happens when a URL is entered in the browser?
1. Browser needs to know how to reach the server (in this example: 'example.com').
2. The process that browser follows in order to know the address of server from domain is called DNS Lookup.
3. DNS: Domain Name System. They are like phonebook of the internet. They translate domain names into IP addresses.
4. To make this process (DNS lookup) fast, the DNS information is heavily cached.
5. Browser itself caches it for a short period of time.
6. Operating System too caches the DNS information for some time, if browser does not find the DNS info in its cache, it asks the OS for it.
7. If OS too does not have it, browser makes a query to the DNS Resolver.
8. This sets up a chain of requests (recursive DNS lookup) until the IP address is resolved.
9. This is an elaborate and elegant process and involves many servers in DNS infrastructures. Also the answer is cached every step of the way.
10. Finally the browser has the IP address of the server 'example.com'.
11. The browser establishes a TCP connection with the server using IP address.
12. The TCP connection establishment involves a TCP handshake process which takes several round trips to complete.
13. To keep the loading process fast, modern browsers use something called 'keep-alive' connection to try to re-use the established connection as mush as possible.
14. This 'keep-alive' looks similar to 'web-sockets'. Where the initial TCP connection is hijaked and a persistent unsupervised connection is established between client and server.
15. Websocket connections are very useful for real-time updates like stock-market apps.
16. If the protocol is 'https', the process of establishment of connection is even more involved. it requires a complicated process of SSL/TLS handshake.
17. The SSL/TLS is an expensive process. So browsers use tricks 'session resumption' to try to lower the cost.
18. Finally the browser sends a request (http or https based) over the established TCP connection.
19. The server processes the request and sends back response.
20. The browser receives the request and renders the response.
21. Oftentimes, there are additional resources to load like js bundles, images.
22. For these additional resources, the browser repeats the chain of action starting from DNS lookup till rendering.
23. So often these static resources are cached onto the browser's cache for quick rendering in subsequent requests.

## Questions remaining to be answered?
1. Details of TCP handshake in case of http and https: [How https work?](https://github.com/HiteshRepo/learning-gist/blob/master/talks/how_https_work.md)
2. What is 'keep-alive' connection? examples, request formats
3. What is 'session resumption'?
4. Difference b/w 'keep-alive' connection and web-sockets?
5. DNS info cache TTLs of browser and OS.
6. DNS lookup process in detail.