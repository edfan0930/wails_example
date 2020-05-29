/* eslint-disable indent */
/* eslint-disable no-alert */
import axios from "axios";
import config from "./config";
import ErrorCode from "./ErrorCode";
import { Message } from "view-design";
import { ToFormData, ToQuerystr, Has } from "../utils";

const axiosIns = axios.create(config);

// 回傳攔截
axiosIns.interceptors.response.use(
  (response) => {
    console.log("reponse", response);
    // 因為使用外部測試網址 所以沒有驗證結構
    // 如果需要驗證註解或刪掉下面一行
    // return response; // 不檢查，直接彈出
    if (!Has(response, "data") || !Has(response.data, "status")) {
      return null;
    }
    const { code } = response.data.status;
    if (code === 0) return response.data.data || true;
    console.log(Message);
    Message.destroy();
    Message.error(ErrorCode(code));
    return null;
  },
  (error) => {
    console.log("in error", Message);
    Message.destroy();
    Message.error(ErrorCode(error.response.status));
    return null;
  }
);

/**
 * Get method
 * @param { String } uri
 * @param { Object } headers
 */
const Get = (uri, Params, headers) => {
  console.log("uri", uri, "params", Params, "headers", headers);
  new Promise((resolve) => {
    axiosIns
      .get(uri + ToQuerystr(Params), {
        headers,
      })
      .then((response) => resolve(response));
  });
};

/**
 * Post method
 * @param { String } uri
 * @param { Object } Params
 * @param { Object } headers
 */
const Post = (uri, Params = {}, headers) =>
  new Promise((resolve) => {
    axiosIns
      .post(uri, ToFormData(Params), {
        headers,
      })
      .then((response) => resolve(response));
  });

/**
 * Put method
 * @param { String } uri
 * @param { Object } Params
 * @param { Object } headers
 */
const Put = (uri, Params = {}, headers) =>
  new Promise((resolve) => {
    axiosIns
      .put(uri, ToFormData(Params), {
        headers,
      })
      .then((response) => resolve(response));
  });

/**
 * Put method
 * @param { String } uri
 * @param { Object } Params
 * @param { Object } headers
 */
const _Put = (uri, Params = {}, headers) =>
  new Promise((resolve) => {
    axiosIns
      .put(uri, Params, {
        headers,
      })
      .then((response) => resolve(response));
  });

/**
 * Delete method
 * @param { String } uri
 * @param { Object } Params
 * @param { Object } headers
 */
const Delete = (uri, Params = {}, headers) =>
  new Promise((resolve) => {
    axiosIns
      .delete(
        uri,
        {
          data: ToFormData(Params),
        },
        {
          headers,
        }
      )
      .then((response) => resolve(response));
  });

export { Get, Post, Put, _Put, Delete };
