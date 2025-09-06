import {useAuthStore} from "/@/store/authStore.ts";


export default {
  mounted: function (el: HTMLElement, binding: any) {
    const authStore = useAuthStore();
    if (authStore && authStore.user && authStore.user.permission) {
      const ps = authStore.user.permission;
      const st = binding.value.split(',')
      for (const s of st) {
        if (ps.includes(s)) {
          return
        }
      }
      if (el != null && el.parentNode != null) {
        el.parentNode.removeChild(el);
      }
    }
  }
}