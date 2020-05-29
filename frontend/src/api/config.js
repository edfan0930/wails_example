export default {
  method: "post",
  baseURL: "http://localhost:9457", // 放上你本地後段IP
  //baseURL: 'https://jsonplaceholder.typicode.com', // 測試用
  transformRequest: [
    function(data, headers) {
      headers.Authorization = localStorage.getItem("token");
      return data;
    },
  ],
  headers: {
    "Content-Type": "multipart/form-data",
    "Access-Control-Allow-Origin": "*",
  },
  // 攜帶憑證
  withCredentials: false,
  // 返回資料型別
  responseType: "json",
};
