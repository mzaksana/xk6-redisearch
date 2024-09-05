<!--

  You can edit the file as you like before or after the HTML comment,
  but do not edit the API documentation between the following HTML comments,
  it was automatically generated from the index.d.ts file.

  You can regenerate the API documentation and bindings code at any time
  by "go generate ." command. The "//go:generate ..." comments required for this
  can be found in the redisearch.go file.

-->
<!-- begin:api -->
xk6-redisearch
==============

A k6 extension to test the performance of a Redisearch

<details><summary><em>Example</em></summary>

```ts
import globalRedisearch, { Redisearch } from "k6/x/redisearch"

export default function () {
  console.log(globalRedisearch.greeting)

  let instance = new Redisearch("Wonderful World")
  console.log(instance.greeting)
}
```

</details>

The [examples](https://github.com/mzaksana/xk6-redisearch/blob/master/examples) directory contains examples of how to use the xk6-redisearch extension. A k6 binary containing the xk6-redisearch extension is required to run the examples. *If the search path also contains the k6 command, don't forget to specify which k6 you want to run (for example `./k6`\)*.

<details>
<summary><strong>Build</strong></summary>

The [xk6](https://github.com/grafana/xk6) build tool can be used to build a k6 that will include xk6-redisearch extension:

```bash
$ xk6 build --with github.com/mzaksana/xk6-redisearch@latest
```

For more build options and how to use xk6, check out the [xk6 documentation]([xk6](https://github.com/grafana/xk6)).

</details>

API
===

Redisearch
----------

This is the primary class of the redisearch extension.

<details><summary><em>Example</em></summary>

```ts
import { Redisearch } from "k6/x/redisearch"

export default function () {
  let instance = new Redisearch("Wonderful World")
  console.log(instance.greeting)
}
```

</details>

### Redisearch()

```ts
constructor(name: string);
```

-	`name` to whom the greeting is addressed

Create a new Redisearch instance.

### Redisearch.greeting

```ts
readonly greeting: string;
```

Greeting message
<!-- end:api -->
