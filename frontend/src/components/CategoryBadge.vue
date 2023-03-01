<template>
  <span class="tag is-light mb-2" :class="color">{{ value | capitalize }}</span>
</template>

<script>
export default {
  props: { value: String },
  computed: {
    color() {
      switch (this.value) {
        case "sport":
          return "is-success";
        case "movies":
          return "is-warning";
        case "finance":
          return "is-danger";
        default:
          return "";
      }
    },
  },
  filters: {
    capitalize(value) {
      if (value === "") {
        return "";
      } else {
        value = value.toString();
        return value.charAt(0).toUpperCase() + value.slice(1);
      }
    },
  },
};
</script>
