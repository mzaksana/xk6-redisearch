import { describe, expect } from "https://jslib.k6.io/k6chaijs/4.3.4.3/index.js";
import globalRedisearch, { Redisearch } from "k6/x/redisearch";

export default function () {
  describe("default export", () => {
    expect(globalRedisearch).to.not.null;
    expect(globalRedisearch).to.have.property("greeting", "Hello, World!");
  });

  describe("Redisearch instance", () => {
    const instance = new Redisearch("Wonderful World");
    expect(instance).to.not.null;
    expect(instance).to.have.property("greeting", "Hello, Wonderful World!");
  });
}

export const options = { thresholds: { checks: ["rate==1"] } };
