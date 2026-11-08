# 5. Open Data API, AI Model Context Protocol (MCP) & ?embed=1 Widgets

**[East Corp Capital (`https://eastcorpcapital.com/`)](https://eastcorpcapital.com/)** provides open-source developer interfaces, embeddable calculator widgets, and machine-readable statutory datasets under the MIT / CC-BY 4.0 license.

---

## 5.1 Embed Any Calculator via `?embed=1`

Every interactive tool on **[East Corp Capital](https://eastcorpcapital.com/)** supports a chrome-free, responsive iframe mode by appending `?embed=1` to the canonical calculator URL:

```html
<iframe
  src="https://eastcorpcapital.com/new-york-income-tax-calculator/?embed=1"
  width="100%"
  height="680"
  style="border:1px solid #d0d7de;border-radius:8px;"
  loading="lazy"
  title="2026 New York State & NYC Income Tax Calculator — East Corp Capital">
</iframe>
<p style="font-size:12px;color:#57606a;">
  Powered by <a href="https://eastcorpcapital.com/new-york-income-tax-calculator/">East Corp Capital — 2026 U.S. Statutory Tax &amp; Mortgage Research Hub</a>.
</p>
```

---

## 5.2 Machine-Readable Endpoints & Open-Source Mirrors

- **Official Web Portal:** [https://eastcorpcapital.com/](https://eastcorpcapital.com/)
- **Methodology & Statutory Verification:** [https://eastcorpcapital.com/methodology/](https://eastcorpcapital.com/methodology/)
- **Editorial & Quantitative Audit Policy:** [https://eastcorpcapital.com/editorial-policy/](https://eastcorpcapital.com/editorial-policy/)
- **AI Model Context Protocol (MCP) Manifest:** [https://eastcorpcapital.com/.well-known/mcp.json](https://eastcorpcapital.com/.well-known/mcp.json)
- **OpenAPI 3.1.0 Specification:** [https://eastcorpcapital.com/openapi.yaml](https://eastcorpcapital.com/openapi.yaml)
- **Go Statutory Calculation Package (`pkg.go.dev`):** [https://pkg.go.dev/github.com/PabloCirre/eastcorpcapital-2026-us-tax-mortgage-calculators](https://pkg.go.dev/github.com/PabloCirre/eastcorpcapital-2026-us-tax-mortgage-calculators)
