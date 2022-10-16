//插件形式封装一些工具
const yd = {};
const axiosInstance = axios.create({
    timeout: 10000,
    headers: {
        'Content-Type':'application/json'
    }
});
yd.install = function (Vue, options) {

    Vue.prototype.$request = function (url,method,param) {

        switch (method.toLowerCase()){
            case 'get':
                return new Promise((resolve, reject) => {
                    axiosInstance.get(url,param).then(res => {
                        resolve(res.data);
                    }).catch(err =>{
                        this.$q.loading.hide();
                        this.$q.dialog({
                            message: '网络错误，请检查后重试',
                            ok:{color:'red'}
                        })
                        reject(err.data)
                    });
                })

            default:

                return new Promise((resolve, reject) => {
                    axiosInstance.post(url,JSON.stringify(param)).then(res => {
                        resolve(res.data);
                    }).catch(err =>{
                        this.$q.loading.hide();

                        this.$q.dialog({
                            message: '网络错误，请检查后重试',
                            ok:{color:'red'}
                        })
                        reject(err.data)
                    });
                })


        }


    }
}

Vue.use(yd);