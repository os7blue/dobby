//插件形式封装一些工具

axios.defaults.withCredentials=true
const axiosInstance = axios.create({
    timeout: 600000,
    headers: {
        'Content-Type': 'application/json'
    }
});

const http = {
    install(app, options) {
        app.config.globalProperties.$request = {
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

                return new Promise((resolve, reject) => {
                    axiosInstance.post(url, JSON.stringify(data)).then(res => {
                        console.log(res)
                        resolve(res.data)
                    }).catch(err => {
                        reject(err.data)
                    })
                })
            }

        }


    }
}