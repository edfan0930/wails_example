<template>
<div>
    <Upload
        multiple
        type="drag"
        :before-upload ="beforeUpload"
        action="//jsonplaceholder.typicode.com/posts/"
    >
        <div style="padding: 20px 0">
            <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
            <p>Click or drag files here to upload</p>
        </div>
    </Upload>
    <Button @click="tryFirst" type="info" />

    <div>
    <video id="myVideo" width="400" autoplay loop="loop">
        <source id="mp4_src" src="" type="video/mp4">
        Your browser does not support HTML video.

    </video>
</div>
</div>
</template>
<script>
    export default {
        name: 'VideoUpload',
        methods: {
            beforeUpload:function(f){
                console.log('19',f);
                console.log('20',f.size)
               const reader = new FileReader()
               reader.onload = f => {
                   this.videoSrc = f.target.result
                   console.log('video src',this.videoSrc)
               }
               //reader.onload = e => this.$emit('load', e.target.result)
                reader.readAsDataURL(f)
                console.log('25',reader)
                return false
            },
            async tryFirst(){
                console.log("in try first");
                let rp = await this.$api.GetFirst();
                this.alertMessage=rp;
                return
            },
        },
        data() {
            return {
                videoSrc:'',
            }
        },
        watch: {
            videoSrc:function(){
               console.log("in computed")
               document.getElementById("mp4_src").src = this.videoSrc;
               document.getElementById("myVideo").load();

           } 
        },
    }
</script>
