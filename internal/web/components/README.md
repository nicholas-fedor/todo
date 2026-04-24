# Todo's TemplUI Components

`components` is used to store components used by the Templ files in the `pages` directory.

Additional components can be added by using the `templui` cli command.
Just remember to ensure that the `.templui.json` configuration points to the directory location.

For example:

```json
{
  "componentsDir": "internal/web/components",
  "utilsDir": "internal/utils",
  "moduleName": "github.com/nicholas-fedor/todo",
  "jsDir": "internal/web/assets/js",
  "jsPublicPath": "/assets/js"
}
```

For example, to add the current set of components, use the following CLI command:

```bash
templui add button card checkbox input label separator
```
