"use strict";(self.webpackChunkui=self.webpackChunkui||[]).push([[540],{2540:(e,r,t)=>{t.r(r),t.d(r,{AlchemyProvider:()=>O});var n=t(5812),o=t(9083),s=t(9042),i=t(5889);t(4523);let c=!1,a=!1;const h={debug:1,default:2,info:2,warning:3,error:4,off:5};let u=h.default,E=null;const l=function(){try{const e=[];if(["NFD","NFC","NFKD","NFKC"].forEach((r=>{try{if("test"!=="test".normalize(r))throw new Error("bad normalize")}catch(t){e.push(r)}})),e.length)throw new Error("missing "+e.join(", "));if(String.fromCharCode(233).normalize("NFD")!==String.fromCharCode(101,769))throw new Error("broken implementation")}catch(e){return e.message}return null}();var N,d;!function(e){e.DEBUG="DEBUG",e.INFO="INFO",e.WARNING="WARNING",e.ERROR="ERROR",e.OFF="OFF"}(N||(N={})),function(e){e.UNKNOWN_ERROR="UNKNOWN_ERROR",e.NOT_IMPLEMENTED="NOT_IMPLEMENTED",e.UNSUPPORTED_OPERATION="UNSUPPORTED_OPERATION",e.NETWORK_ERROR="NETWORK_ERROR",e.SERVER_ERROR="SERVER_ERROR",e.TIMEOUT="TIMEOUT",e.BUFFER_OVERRUN="BUFFER_OVERRUN",e.NUMERIC_FAULT="NUMERIC_FAULT",e.MISSING_NEW="MISSING_NEW",e.INVALID_ARGUMENT="INVALID_ARGUMENT",e.MISSING_ARGUMENT="MISSING_ARGUMENT",e.UNEXPECTED_ARGUMENT="UNEXPECTED_ARGUMENT",e.CALL_EXCEPTION="CALL_EXCEPTION",e.INSUFFICIENT_FUNDS="INSUFFICIENT_FUNDS",e.NONCE_EXPIRED="NONCE_EXPIRED",e.REPLACEMENT_UNDERPRICED="REPLACEMENT_UNDERPRICED",e.UNPREDICTABLE_GAS_LIMIT="UNPREDICTABLE_GAS_LIMIT",e.TRANSACTION_REPLACED="TRANSACTION_REPLACED",e.ACTION_REJECTED="ACTION_REJECTED"}(d||(d={}));const g="0123456789abcdef";class R{constructor(e){Object.defineProperty(this,"version",{enumerable:!0,value:e,writable:!1})}_log(e,r){const t=e.toLowerCase();null==h[t]&&this.throwArgumentError("invalid log level name","logLevel",e),u>h[t]||console.log.apply(console,r)}debug(){for(var e=arguments.length,r=new Array(e),t=0;t<e;t++)r[t]=arguments[t];this._log(R.levels.DEBUG,r)}info(){for(var e=arguments.length,r=new Array(e),t=0;t<e;t++)r[t]=arguments[t];this._log(R.levels.INFO,r)}warn(){for(var e=arguments.length,r=new Array(e),t=0;t<e;t++)r[t]=arguments[t];this._log(R.levels.WARNING,r)}makeError(e,r,t){if(a)return this.makeError("censored error",r,{});r||(r=R.errors.UNKNOWN_ERROR),t||(t={});const n=[];Object.keys(t).forEach((e=>{const r=t[e];try{if(r instanceof Uint8Array){let t="";for(let e=0;e<r.length;e++)t+=g[r[e]>>4],t+=g[15&r[e]];n.push(e+"=Uint8Array(0x"+t+")")}else n.push(e+"="+JSON.stringify(r))}catch(i){n.push(e+"="+JSON.stringify(t[e].toString()))}})),n.push(`code=${r}`),n.push(`version=${this.version}`);const o=e;let s="";switch(r){case d.NUMERIC_FAULT:{s="NUMERIC_FAULT";const r=e;switch(r){case"overflow":case"underflow":case"division-by-zero":s+="-"+r;break;case"negative-power":case"negative-width":s+="-unsupported";break;case"unbound-bitwise-result":s+="-unbound-result"}break}case d.CALL_EXCEPTION:case d.INSUFFICIENT_FUNDS:case d.MISSING_NEW:case d.NONCE_EXPIRED:case d.REPLACEMENT_UNDERPRICED:case d.TRANSACTION_REPLACED:case d.UNPREDICTABLE_GAS_LIMIT:s=r}s&&(e+=" [ See: https://links.ethers.org/v5-errors-"+s+" ]"),n.length&&(e+=" ("+n.join(", ")+")");const i=new Error(e);return i.reason=o,i.code=r,Object.keys(t).forEach((function(e){i[e]=t[e]})),i}throwError(e,r,t){throw this.makeError(e,r,t)}throwArgumentError(e,r,t){return this.throwError(e,R.errors.INVALID_ARGUMENT,{argument:r,value:t})}assert(e,r,t,n){e||this.throwError(r,t,n)}assertArgument(e,r,t,n){e||this.throwArgumentError(r,t,n)}checkNormalize(e){l&&this.throwError("platform missing String.prototype.normalize",R.errors.UNSUPPORTED_OPERATION,{operation:"String.prototype.normalize",form:l})}checkSafeUint53(e,r){"number"===typeof e&&(null==r&&(r="value not safe"),(e<0||e>=9007199254740991)&&this.throwError(r,R.errors.NUMERIC_FAULT,{operation:"checkSafeInteger",fault:"out-of-safe-range",value:e}),e%1&&this.throwError(r,R.errors.NUMERIC_FAULT,{operation:"checkSafeInteger",fault:"non-integer",value:e}))}checkArgumentCount(e,r,t){t=t?": "+t:"",e<r&&this.throwError("missing argument"+t,R.errors.MISSING_ARGUMENT,{count:e,expectedCount:r}),e>r&&this.throwError("too many arguments"+t,R.errors.UNEXPECTED_ARGUMENT,{count:e,expectedCount:r})}checkNew(e,r){e!==Object&&null!=e||this.throwError("missing new",R.errors.MISSING_NEW,{name:r.name})}checkAbstract(e,r){e===r?this.throwError("cannot instantiate abstract class "+JSON.stringify(r.name)+" directly; use a sub-class",R.errors.UNSUPPORTED_OPERATION,{name:e.name,operation:"new"}):e!==Object&&null!=e||this.throwError("missing new",R.errors.MISSING_NEW,{name:r.name})}static globalLogger(){return E||(E=new R("logger/5.7.0")),E}static setCensorship(e,r){if(!e&&r&&this.globalLogger().throwError("cannot permanently disable censorship",R.errors.UNSUPPORTED_OPERATION,{operation:"setCensorship"}),c){if(!e)return;this.globalLogger().throwError("error censorship permanent",R.errors.UNSUPPORTED_OPERATION,{operation:"setCensorship"})}a=!!e,c=!!r}static setLogLevel(e){const r=h[e.toLowerCase()];null!=r?u=r:R.globalLogger().warn("invalid log level - "+e)}static from(e){return new R(e)}}R.errors=d,R.levels=N;const f=new R("properties/5.7.0");function p(e,r,t){Object.defineProperty(e,r,{enumerable:!0,value:t,writable:!1})}const m={bigint:!0,boolean:!0,function:!0,number:!0,string:!0};function I(e){if(void 0===e||null===e||m[typeof e])return!0;if(Array.isArray(e)||"object"===typeof e){if(!Object.isFrozen(e))return!1;const t=Object.keys(e);for(let n=0;n<t.length;n++){let o=null;try{o=e[t[n]]}catch(r){continue}if(!I(o))return!1}return!0}return f.throwArgumentError("Cannot deepCopy "+typeof e,"object",e)}function w(e){if(I(e))return e;if(Array.isArray(e))return Object.freeze(e.map((e=>A(e))));if("object"===typeof e){const r={};for(const t in e){const n=e[t];void 0!==n&&p(r,t,A(n))}return r}return f.throwArgumentError("Cannot deepCopy "+typeof e,"object",e)}function A(e){return w(e)}class _{constructor(e){let r=arguments.length>1&&void 0!==arguments[1]?arguments[1]:100;this.sendBatchFn=e,this.maxBatchSize=r,this.pendingBatch=[]}enqueueRequest(e){return(0,n._)(this,void 0,void 0,(function*(){const r={request:e,resolve:void 0,reject:void 0},t=new Promise(((e,t)=>{r.resolve=e,r.reject=t}));return this.pendingBatch.push(r),this.pendingBatch.length===this.maxBatchSize?this.sendBatchRequest():this.pendingBatchTimer||(this.pendingBatchTimer=setTimeout((()=>this.sendBatchRequest()),10)),t}))}sendBatchRequest(){return(0,n._)(this,void 0,void 0,(function*(){const e=this.pendingBatch;this.pendingBatch=[],this.pendingBatchTimer&&(clearTimeout(this.pendingBatchTimer),this.pendingBatchTimer=void 0);const r=e.map((e=>e.request));return this.sendBatchFn(r).then((r=>{e.forEach(((e,t)=>{const n=r[t];if(n.error){const r=new Error(n.error.message);r.code=n.error.code,r.data=n.error.data,e.reject(r)}else e.resolve(n.result)}))}),(r=>{e.forEach((e=>{e.reject(r)}))}))}))}}class O extends s.r{constructor(e){const r=O.getApiKey(e.apiKey),t=O.getAlchemyNetwork(e.network),o=O.getAlchemyConnectionInfo(t,r,"http");void 0!==e.url&&(o.url=e.url),o.throttleLimit=e.maxRetries;super(o,n.E[t]),this.apiKey=e.apiKey,this.maxRetries=e.maxRetries,this.batchRequests=e.batchRequests;const s=Object.assign({},this.connection);s.headers["Alchemy-Ethers-Sdk-Method"]="batchSend";this.batcher=new _((e=>(0,i.rd)(s,JSON.stringify(e))))}static getApiKey(e){if(null==e)return n.D;if(e&&"string"!==typeof e)throw new Error(`Invalid apiKey '${e}' provided. apiKey must be a string.`);return e}static getNetwork(e){return"string"===typeof e&&e in n.C?n.C[e]:(0,o.H)(e)}static getAlchemyNetwork(e){if(void 0===e)return n.a;if("number"===typeof e)throw new Error(`Invalid network '${e}' provided. Network must be a string.`);if(!Object.values(n.N).includes(e))throw new Error(`Invalid network '${e}' provided. Network must be one of: ${Object.values(n.N).join(", ")}.`);return e}static getAlchemyConnectionInfo(e,r,t){const o="http"===t?(0,n.g)(e,r):(0,n.b)(e,r);return{headers:n.I?{"Alchemy-Ethers-Sdk-Version":n.V}:{"Alchemy-Ethers-Sdk-Version":n.V,"Accept-Encoding":"gzip"},allowGzip:!0,url:o}}detectNetwork(){const e=Object.create(null,{detectNetwork:{get:()=>super.detectNetwork}});return(0,n._)(this,void 0,void 0,(function*(){let r=this.network;if(null==r&&(r=yield e.detectNetwork.call(this),!r))throw new Error("No network detected");return r}))}_startPending(){(0,n.l)("WARNING: Alchemy Provider does not support pending filters")}isCommunityResource(){return this.apiKey===n.D}send(e,r){return this._send(e,r,"send")}_send(e,r,t){let n=arguments.length>3&&void 0!==arguments[3]&&arguments[3];const o={method:e,params:r,id:this._nextId++,jsonrpc:"2.0"};if(Object.assign({},this.connection).headers["Alchemy-Ethers-Sdk-Method"]=t,this.batchRequests||n)return this.batcher.enqueueRequest(o);this.emit("debug",{action:"request",request:A(o),provider:this});const s=["eth_chainId","eth_blockNumber"].indexOf(e)>=0;if(s&&this._cache[e])return this._cache[e];const c=(0,i.rd)(this.connection,JSON.stringify(o),T).then((e=>(this.emit("debug",{action:"response",request:o,response:e,provider:this}),e)),(e=>{throw this.emit("debug",{action:"response",error:e,request:o,provider:this}),e}));return s&&(this._cache[e]=c,setTimeout((()=>{this._cache[e]=null}),0)),c}}function T(e){if(e.error){const r=new Error(e.error.message);throw r.code=e.error.code,r.data=e.error.data,r}return e.result}}}]);
//# sourceMappingURL=540.b38eb83e.chunk.js.map