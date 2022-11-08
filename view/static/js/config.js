//插件形式封装一些工具

axios.defaults.withCredentials = true
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
                        resolve(res.data)
                    }).catch(err => {
                        reject(err.data)
                    })
                })
            }

        }


    }
}


function pushZero(num) {
    return num < 10 ? (`0${  num}`) : num
}
const util = {
    install(app, options) {
        app.config.globalProperties.$util = {

            longTimeFormat: (longTime) => {
                if (!longTime) {
                    return ""
                }

                const time = new Date(Number(longTime))
                return `${time.getFullYear()}-${pushZero(time.getMonth() + 1)}-${pushZero(time.getDate())} ${pushZero(time.getHours())}:${pushZero(time.getMinutes())}:${pushZero(time.getSeconds())}`
            },
            longTimeWeekFormat: (longTime) => {
                if (!longTime) {
                    return ""
                }
                const time = new Date(Number(longTime))
                const weeks = ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
                return `${time.getFullYear()}-${time.getMonth() + 1}-${time.getDate()} ${time.getHours()}:${time
                    .getMinutes()}:${time.getSeconds()} ${weeks[time.getDay()]}`
            },
            arrayParamToStr: (key = '', data = []) => {

                let params = '';
                if (key === '') {
                    for (let i = 0; i < data.length; i++) {


                        const param = data[i];
                        if (i !== 0 && data.length > 1) {
                            params += `,${param}`

                        } else {
                            params += param;

                        }

                    }
                    return params;

                }

                for (let i = 0; i < data.length; i++) {


                    const param = data[i][key];
                    if (i !== 0 && data.length > 1) {
                        params += `,${param}`

                    } else {
                        params += param;

                    }
                }
                return params;


            },


            getArrayItemIndex: (arr, item, key) => {

                for (let i = 0; i < arr.length; i++) {
                    let temp = arr[i]

                    if (temp[key] === item[key]) {
                        return i
                    }

                }

                return -1

            }

        }


    }
}