import http from "k6/http";
import { check } from "k6";

export default function() {
  let params = {
    timeout: 10000
  };
  let res = http.get("http://localhost:8081/dobusiness", params);
  check(res, {
      "response status 200": (r) => r.status === 200
  });
};
