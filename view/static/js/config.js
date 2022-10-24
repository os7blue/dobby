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

const util = {
    install(app, options) {
        app.config.globalProperties.$util = {
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