const glob = require("glob");
const markdownIt = require("markdown-it");
const meta = require("markdown-it-meta");
const fs = require("fs");
const _ = require("lodash");

const sidebar = (directory, array) => {
    return array.map(i => {
        const children = _.sortBy(
            glob
                .sync(`./${directory}/${i[1]}/*.md`)
                .map(path => {
                    const md = new markdownIt();
                    const file = fs.readFileSync(path, "utf8");
                    md.use(meta);
                    md.render(file);
                    const order = md.meta.order;
                    return { path, order };
                })
                .filter(f => f.order !== false),
            ["order", "path"]
        )
            .map(f => f.path)
            .filter(f => !f.match("README"));

        return {
            title: i[0],
            children
        };
    });
};

module.exports = {
    head: [
        ['link', { rel: 'icon', href: '/logo.png' }],
        ['link', { rel: 'shortcut icon', type: "image/x-icon", href: "/logo.png" }],
        ['meta', { name: 'theme-color', content: '#14e6c7' }],
    ],
    locales: {
        "/": {
            lang: "en-US",
            title: "BHP Documents",
            description: "BHP Documents",
        },
        "/zh/": {
            lang: "简体中文",
            title: "BHP 文档",
            description: "BHP 文档",
        }
    },
    plugins: [['vuepress-plugin-code-copy', true]],
    base:'/',
    themeConfig: {
        sidebarDepth: 0,
        locales: {
            '/': {
                selectText: 'Languages',
                label: 'English',
                ariaLabel: 'Languages',
                editLinkText: 'Edit this page on GitHub',
                serviceWorker: {
                    updatePopup: {
                        message: "New content is available.",
                        buttonText: "Refresh"
                    }
                },
                algolia: {},
                nav: [
                    { text: 'Back to BHP', link: 'https://bhpnet.io' },
                    { text: 'BHP Github', link: 'https://github.com/bhpnet' },
                    { text: 'BHP Faucet', link: 'https://faucet.bhpnet.io/' },
                    { text: 'BHP Explorer', link: 'https://scan.bhpnet.io' }
                ],
                sidebar:sidebar("", [
                    ["BHP", "/bhp"],
                    ["Getting Start", "getting-start"],
                    ["Delegators", "delegators"],
                    ["Validators", "validators"],
                    ["Concepts", "concepts"],
                    ["Daemon", "daemon"],
                    ["CLI Client", "cli-client"],
                    ["Resources", "resources"],
                ])
            },
            '/zh/': {
                // 多语言下拉菜单的标题
                selectText: '选择语言',
                // 该语言在下拉菜单中的标签
                label: '简体中文',
                // 编辑链接文字
                editLinkText: '在 GitHub 上编辑此页',
                // Service Worker 的配置
                serviceWorker: {
                    updatePopup: {
                        message: "发现新内容可用.",
                        buttonText: "刷新"
                    }
                },
                // 当前 locale 的 algolia docsearch 选项
                algolia: {},
                nav: [
                    { text: 'BHP 官网', link: 'https://bhpnet.io' },
                    { text: 'BHP Github', link: 'https://github.com/bhpnet' },
                    { text: 'BHP 水龙头', link: 'https://faucet.bhpnet.io/' },
                    { text: 'BHP 浏览器', link: 'https://scan.bhpnet.io' }
                ],
                sidebar: sidebar("", [
                    ["BHP简介", "/zh/bhp"],
                    ["快速入门", "/zh/getting-start"],
                    ["委托者", "/zh/delegators"],
                    ["验证者", "/zh/validators"],
                    ["概念", "/zh/concepts"],
                    ["守护进程", "/zh/daemon"],
                    ["命令行客户端", "/zh/cli-client"],
                    ["工具", "/zh/tools"],
                    ["资源", "/zh/resources"],
                ])
            }
        }
    }
}