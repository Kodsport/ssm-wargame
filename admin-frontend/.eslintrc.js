module.exports = {
  root: true,
  parser: "vue-eslint-parser",
  parserOptions: {
    parser: "@typescript-eslint/parser",
  },
  env: {
    browser: true,
  },
  extends: [
    "plugin:@typescript-eslint/recommended",
    "plugin:vue/recommended",
  //  "@vue/prettier",
    "plugin:prettier/recommended",
    "@vue/typescript",
  ],
  // required to lint *.vue files
  plugins: ["vue", "@typescript-eslint"],
  // add your custom rules here
  rules: {
    // allow async-await
    "generator-star-spacing": "off",
    // allow debugger during development
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-console": process.env.NODE_ENV === "production" ? "error" : "off",
  },
};
