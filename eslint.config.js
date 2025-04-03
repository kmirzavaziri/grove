const tsEslint = require('@typescript-eslint/eslint-plugin');
const importPlugin = require('eslint-plugin-import');

module.exports = [
    {
        files: ['ts/**/*.{ts,tsx}'],
        languageOptions: {
            parser: require('@typescript-eslint/parser'),
            parserOptions: {
                project: [
                    'ts/packages/grove/tsconfig.json',
                    'ts/packages/grovex/tsconfig.json',
                    'ts/examples/dashboard/tsconfig.json',
                ]
            }
        },
        plugins: {
            '@typescript-eslint': tsEslint,
            import: importPlugin
        },
        rules: {
            'import/no-unresolved': 'error',
            '@typescript-eslint/no-unused-vars': 'warn',
            'import/order': 'warn'
        },
        settings: {
            'import/resolver': {
                typescript: {
                    alwaysTryTypes: true,
                    "project": [
                        'ts/packages/grove/tsconfig.json',
                        'ts/packages/grovex/tsconfig.json',
                        "ts/examples/dashboard/tsconfig.json"
                    ]
                }
            }
        }
    }
];