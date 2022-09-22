"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[4356],{3905:(e,t,r)=>{r.d(t,{Zo:()=>l,kt:()=>f});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var o=n.createContext({}),u=function(e){var t=n.useContext(o),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},l=function(e){var t=u(e.components);return n.createElement(o.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},b=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,s=e.originalType,o=e.parentName,l=c(e,["components","mdxType","originalType","parentName"]),b=u(r),f=a,k=b["".concat(o,".").concat(f)]||b[f]||p[f]||s;return r?n.createElement(k,i(i({ref:t},l),{},{components:r})):n.createElement(k,i({ref:t},l))}));function f(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var s=r.length,i=new Array(s);i[0]=b;var c={};for(var o in t)hasOwnProperty.call(t,o)&&(c[o]=t[o]);c.originalType=e,c.mdxType="string"==typeof e?e:a,i[1]=c;for(var u=2;u<s;u++)i[u]=r[u];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}b.displayName="MDXCreateElement"},4182:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>o,contentTitle:()=>i,default:()=>p,frontMatter:()=>s,metadata:()=>c,toc:()=>u});var n=r(7462),a=(r(7294),r(3905));const s={},i=void 0,c={unversionedId:"cli-reference/kubectl-testkube_create_testsuite",id:"cli-reference/kubectl-testkube_create_testsuite",title:"kubectl-testkube_create_testsuite",description:"kubectl-testkube create testsuite",source:"@site/docs/5-cli-reference/kubectl-testkube_create_testsuite.md",sourceDirName:"5-cli-reference",slug:"/cli-reference/kubectl-testkube_create_testsuite",permalink:"/testkube/cli-reference/kubectl-testkube_create_testsuite",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/5-cli-reference/kubectl-testkube_create_testsuite.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"kubectl-testkube_create_testsource",permalink:"/testkube/cli-reference/kubectl-testkube_create_testsource"},next:{title:"kubectl-testkube_create_ticket",permalink:"/testkube/cli-reference/kubectl-testkube_create_ticket"}},o={},u=[{value:"kubectl-testkube create testsuite",id:"kubectl-testkube-create-testsuite",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],l={toc:u};function p(e){let{components:t,...r}=e;return(0,a.kt)("wrapper",(0,n.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"kubectl-testkube-create-testsuite"},"kubectl-testkube create testsuite"),(0,a.kt)("p",null,"Create new TestSuite"),(0,a.kt)("h3",{id:"synopsis"},"Synopsis"),(0,a.kt)("p",null,"Create new TestSuite Custom Resource"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"kubectl-testkube create testsuite [flags]\n")),(0,a.kt)("h3",{id:"options"},"Options"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"      --execution-name string                      execution name, if empty will be autogenerated\n  -f, --file string                                JSON test suite file - will be read from stdin if not specified, look at testkube.TestUpsertRequest\n  -h, --help                                       help for testsuite\n      --http-proxy string                          http proxy for executor containers\n      --https-proxy string                         https proxy for executor containers\n  -l, --label stringToString                       label key value pair: --label key1=value1 (default [])\n      --name string                                Set/Override test suite name\n      --schedule string                            test suite schedule in a cronjob form: * * * * *\n  -s, --secret-variable stringToString             secret variable key value pair: --secret-variable key1=value1 (default [])\n      --secret-variable-reference stringToString   secret variable references in a form name1=secret_name1=secret_key1 (default [])\n  -v, --variable stringToString                    param key value pair: --variable key1=value1 (default [])\n")),(0,a.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'  -a, --api-uri string     api uri, default value read from config if set (default "http://localhost:8088")\n  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")\n      --crd-only           generate only crd\n      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")\n      --oauth-enabled      enable oauth\n      --verbose            show additional debug messages\n')),(0,a.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/testkube/cli-reference/kubectl-testkube_create"},"kubectl-testkube create"),"\t - Create resource")))}p.isMDXComponent=!0}}]);