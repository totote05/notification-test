<template>
  <span class="tag mb-1" :class="color">{{ value }}</span>
</template>

<script>
export default {
  props: { value: String },
  computed: {
    color() {
      switch (this.value) {
        case "email":
          return "is-primary";
        case "sms":
          return "is-link";
        case "push":
          return "is-warning";
        default:
          return "is-light";
      }
    },
  },
};
</script>
