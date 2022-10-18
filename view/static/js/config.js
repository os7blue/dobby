//插件形式封装一些工具
const yd = {};
const axiosInstance = axios.create({
    timeout: 600000,
    headers: {
        'Content-Type': 'application/json'
    }
});
yd.install = function (Vue, options) {

    Vue.prototype.$request = {
        get: (url = '', data = {}) => {
            return new Promise((resolve, reject) => {
                axiosInstance.get(url, JSON.stringify(data)).then(res => {
                    resolve(res.data)
                }).catch(err => {
                    reject(err.data)
                })
            })
        },
        post: (url = '', data = {}) => {
            console.log(data)
            return new Promise((resolve, reject) => {
                axiosInstance.post(url, JSON.stringify(data)).then(res => {
                    resolve(res.data)
                }).catch(err => {
                    reject(err.data)
                })
            })
        }

    }


}

Vue.use(yd);