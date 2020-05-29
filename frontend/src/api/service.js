import { Get } from "./instance";
import router from "./router";

export default {
  /**
   * Login
   * @param { Object } Obj
   * @property { account }
   * @property { password }
   */
  Login: (Obj) => Get(router.LOGIN, Obj),
  CounterGet: (cid) => Get(router.COUNTER.concat("/" + cid)),
  CounterNext: (cid) => Get(router.COUNTER.concat("/next/" + cid)),
  GoVersion: () => Get("/9527/version"),
  GetFirst: () => Get("/first"),
};
