//插件形式封装一些工具
const p = {};
const axiosInstance = axios.create({
    timeout: 10000,
    headers: {
        'Content-Type':'application/json'
    }
});
p.install = function (Vue, options) {

    Vue.prototype.$request = function (url,method,param) {

        switch (method.toLowerCase()){
            case 'get':
                return new Promise((resolve, reject) => {
                    axiosInstance.get(url,param).then(res => {
                        resolve(res.data);
                    }).catch(err =>{

                        reject(err.data)
                    });
                })
            default:

                return new Promise((resolve, reject) => {
                    axiosInstance.post(url,param).then(res => {
                        resolve(res.data);
                    }).catch(err =>{

                        reject(err.data)
                    });
                })


        }


    }
}

Vue.use(p);