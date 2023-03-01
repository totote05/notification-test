<template>
  <span class="tag" :class="status.color">{{ status.label }}</span>
</template>

<script>
export default {
  props: { value: String },
  computed: {
    status() {
      if (this.value === "response_ok") {
        return { label: "OK", color: "is-success" };
      } else {
        return { label: "Fail", color: "is-danger" };
      }
    },
  },
};
</script>
