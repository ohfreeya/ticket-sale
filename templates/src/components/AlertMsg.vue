<template>
    <div>
        <v-alert
            :msg = msg
            :type = type
            :value = isVisable
        >
        {{ msg }}
        </v-alert>
    </div>
</template>
<script>
import { ref } from 'vue'
export default {
    name: 'AlertMsg',
    props: {
        msg: {
            type: String,
            default: ''
        },
        type: {
            type: String,
            default: 'info'
        },
        isVisable: {
            type: Boolean,
            default: false
        }
    },
    setup(props) {
        const msg1 = ref(props.msg)
        const type1 = ref(props.type)
        const isVisable1 = ref(props.isVisable)

        setTimeout(() => {
            isVisable1.value = false
        }, 3000)
        return {
            msg1,
            type1,
            isVisable1
        };
    }
}
</script>