import globalRedisearch, { Redisearch } from "k6/x/redisearch";

export default function () {
  console.log(globalRedisearch.greeting);

  let instance = new Redisearch("Wonderful World");
  console.log(instance.greeting);
}
