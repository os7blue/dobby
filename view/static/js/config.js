//插件形式封装一些工具

axios.defaults.withCredentials = true
const axiosInstance = axios.create({
    timeout: 600000,
    headers: {
        'Content-Type': 'application/json'
    }
});
axiosInstance.interceptors.response.use(function (response) {

    if (response.data.code === 401) {
        let index = LayuiVue.layer.confirm(`您的登录状态已失效，点击跳转重新登陆`, {
            icon: 3,
            title: false,
            yes: () => {
                window.location.href = "/"
            }
        })
    }


    return response;
}, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
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
    return num < 10 ? (`0${num}`) : num
}

const util = {
    install(app, options) {
        app.config.globalProperties.$util = {

            /*这个方法用来随机一个十六进制颜色代码，让每一次点击浮动文字的杨色不同*/
            randomHexColor: () => {
                const colorArr = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f']

                let color = "#";
                for (let i = 0; i < 6; i++) {
                    color += colorArr[Math.floor(Math.random() * 16)];
                }
                return color;
            },
            arrayItemParamToArray: (key = '', data = []) => {
                let arr = []
                for (let i = 0; i < data.length; i++) {
                    if (key === '') {
                        arr.push(data[i])
                    } else {
                        arr.push(data[i][key])
                    }
                }

                return arr

            },

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
                return `${time.getFullYear()}-${pushZero(time.getMonth() + 1)}-${pushZero(time.getDate())} ${pushZero(time.getHours())}:${pushZero(time.getMinutes())}:${pushZero(time.getSeconds())} ${weeks[time.getDay()]}`
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