import http from "k6/http";
import { check } from "k6";

export default function() {
  let res = http.get("http://localhost:8081/dobusiness");
  check(res, {
      "response status 200": (r) => r.status === 200
  });
};
