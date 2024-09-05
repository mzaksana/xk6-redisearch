/**
 * A k6 extension to test the performance of a Redisearch
 *
 * @example
 * ```ts
 * import globalRedisearch, { Redisearch } from "k6/x/redisearch"
 *
 * export default function () {
 *   console.log(globalRedisearch.greeting)
 *
 *   let instance = new Redisearch("Wonderful World")
 *   console.log(instance.greeting)
 * }
 * ```
 */
export as namespace redisearch;

/**
 * This is the primary class of the redisearch extension.
 *
 * @example
 * ```ts
 * import { Redisearch } from "k6/x/redisearch"
 *
 * export default function () {
 *   let instance = new Redisearch("Wonderful World")
 *   console.log(instance.greeting)
 * }
 * ```
 */
export declare class Redisearch {
  /** Greeting message */
  readonly greeting: string;

  /**
   * Create a new Redisearch instance.
   *
   * @param name to whom the greeting is addressed
   */
  constructor(name: string);
}

/** Default Redisearch instance. */
declare const defaultRedisearch: Redisearch;

/** Default Redisearch instance. */
export default defaultRedisearch;
