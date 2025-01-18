// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Base() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html class=\"no-js\" lang=\"\"><head><meta charset=\"utf-8\"><title>j-yw dev portfolio</title><link rel=\"shortcut icon\" href=\"/favicon.ico\"><link rel=\"apple-touch-icon\" href=\"icon.png\"><link rel=\"stylesheet\" href=\"/assets/css/normalize.css\"><link rel=\"stylesheet\" href=\"/assets/css/style.css\"><script async src=\"https://www.googletagmanager.com/gtag/js?id=G-BSBD6MP3H2\"></script><script>\n\t\twindow.dataLayer = window.dataLayer || [];\n\t\tfunction gtag() {\n\t\t\tdataLayer.push(arguments);\n\t\t}\n\t\tgtag(\"js\", new Date());\n\n\t\tgtag(\"config\", \"G-BSBD6MP3H2\");\n\t</script></head><body><!-- Add your site or application content here --><header class=\"header\"><h1>Hello!👋🏼</h1></header><main><h2>I'm a frontend-focused fullstack developer with 5 years of experience in crafting high-quality software. I specialize in TypeScript and Golang, and I’ve worked with blockchain platforms like Ethereum and Polygon. Familiar with tools like Solidity, Ethers.js, and Hardhat, I'm always excited to dive into new technologies.</h2><br><h2>You can check out my <a href=\"/assets/media/j-yw_resume.pdf\">Resume</a> here.</h2><br><section class=\"projects\"><h2>Projects</h2><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://cosmology.zone/\" target=\"_blank\" class=\"thumb img-9\"></a><article><h1>Cosmology</h1><p>Build web3 apps in light speed Cosmology develops cutting-edge tools for the interchain ecosystem, empowering seamless interactions across the Internet of Blockchains.</p><span>j-yw 2024 <a href=\"https://cosmology.zone/\" target=\"_blank\">Website</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://www.adsrsounds.com/\" target=\"_blank\" class=\"thumb img-8\"></a><article><h1>ADSR Sounds</h1><p>ADSR Sample Packs & Loops, Synth Presets, Plug-ins & Video Courses for Electronic Music Producers</p><span>j-yw 2023 <a href=\"https://www.adsrsounds.com/\" target=\"_blank\">Website</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://coinvise.co/\" target=\"_blank\" class=\"thumb img-6\"></a><article><h1>Coinvise</h1><p>Powerful Tools For Creators To Build & Operate Web3 Communities</p><span>j-yw 2022 <a href=\"https://coinvise.co/\" target=\"_blank\">Website</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://www.coindesk.com/learn/what-is-desk-coindesks-social-token-explained/ \" target=\"_blank\" class=\"thumb img-7\"></a><article><h1>DESK by Coindesk</h1><p>DESK is CoinDesk’s social token.</p><span>j-yw 2022 <a href=\"https://github.com/j-yw/\" target=\"_blank\">Github</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://steep.surf\" target=\"_blank\" class=\"thumb img-5\"></a><article><h1>Steep Surf Reports</h1><p>Get weather reports for surfing base on location or coordinates</p><span>j-yw 2021 <a href=\"https://github.com/j-yw/steep-surf\" target=\"_blank\">Github</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://nclone-reactjs.netlify.app\" target=\"_blank\" class=\"thumb img-1\"></a><article><h1>Netflix</h1><p>A Netflix Clone built using React, Styled-Components with Firebase database and authentication.</p><span>j-yw 2020 <a href=\"https://github.com/j-yw/netflix\" target=\"_blank\">Github</a></span></article></div></div><div class=\"card-wrapper\"><div class=\"card\"><a href=\"https://todoist-clonexyz.netlify.app\" target=\"_blank\" class=\"thumb img-2\"></a><article><h1>Todoist</h1><p>Fully functional Todoist app built with React, using Firebase as the database and authentication</p><span>j-yw 2020 <a href=\"https://github.com/j-yw/todoist\" target=\"_blank\">Github</a></span></article></div></div></section></main><footer><h1>Thank you! 👏🏼</h1></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
