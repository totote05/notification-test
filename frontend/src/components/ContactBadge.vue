<template>
  <span class="tag" :class="contact.color">{{ contact.label }}</span>
</template>

<script>
export default {
  props: { value: Object },
  computed: {
    contact() {
      switch (this.value.response.channel) {
        case "email":
          return {
            label: this.value.user.email,
            color: "is-primary",
          };
        case "sms":
          return {
            label: this.value.user.phoneNumber,
            color: "is-link",
          };
        case "push":
          return {
            label: this.value.user.phoneNumber,
            color: "is-warning",
          };
        default:
          return {
            label: "-",
            color: "is-light",
          };
      }
    },
  },
};
</script>
