#!/usr/bin/env python3
"""Batch-generate the 30 deck slides via OpenAI gpt-image-2.

Reads SLIDE-DECK-30-VI.md, extracts the fenced prompt under each
"## SLIDE N — ..." header, and renders one PNG per slide.

Requires:  pip install openai
           export OPENAI_API_KEY=sk-...   (needs API billing; ChatGPT login does NOT work here)

Usage:     python3 gen_slides.py
           python3 gen_slides.py 1 5 12     # only specific slide numbers
"""
import os, re, sys, base64, pathlib, time

MD   = pathlib.Path(__file__).with_name("SLIDE-DECK-30-VI.md")
OUT  = pathlib.Path(__file__).with_name("slides-out")
MODEL = "gpt-image-2"
SIZE  = "1536x1024"   # landscape ~3:2 (closest gpt-image-2 size to 16:9)
QUALITY = "high"

def extract_prompts(md_text):
    """Return {slide_number: prompt_text}."""
    prompts = {}
    # split on slide headers, keep the number
    parts = re.split(r'^## SLIDE (\d+)\b.*$', md_text, flags=re.M)
    # parts = [pre, "1", body1, "2", body2, ...]
    for i in range(1, len(parts), 2):
        num  = int(parts[i])
        body = parts[i + 1]
        m = re.search(r'```(?:\w*)\n(.*?)```', body, flags=re.S)
        if m:
            prompts[num] = m.group(1).strip()
    return prompts

def main():
    try:
        from openai import OpenAI
    except ImportError:
        sys.exit("Missing dependency. Run:  pip install openai")

    if not os.environ.get("OPENAI_API_KEY"):
        sys.exit("OPENAI_API_KEY not set. Need an OpenAI API key WITH BILLING.\n"
                 "  export OPENAI_API_KEY=sk-...")

    client = OpenAI()
    OUT.mkdir(exist_ok=True)
    prompts = extract_prompts(MD.read_text(encoding="utf-8"))
    if len(prompts) != 30:
        print(f"WARN: parsed {len(prompts)} prompts (expected 30): {sorted(prompts)}")

    wanted = [int(a) for a in sys.argv[1:]] or sorted(prompts)
    for n in wanted:
        if n not in prompts:
            print(f"slide {n}: no prompt found, skip"); continue
        dest = OUT / f"slide-{n:02d}.png"
        if dest.exists():
            print(f"slide {n:02d}: exists, skip"); continue
        print(f"slide {n:02d}: generating ...", flush=True)
        for attempt in range(1, 4):
            try:
                r = client.images.generate(
                    model=MODEL, prompt=prompts[n], size=SIZE, quality=QUALITY)
                dest.write_bytes(base64.b64decode(r.data[0].b64_json))
                print(f"slide {n:02d}: saved -> {dest}")
                break
            except Exception as e:
                print(f"slide {n:02d}: attempt {attempt} failed: {e}")
                if attempt == 3:
                    print(f"slide {n:02d}: GIVING UP")
                else:
                    time.sleep(5 * attempt)
    print("done. files in", OUT)

if __name__ == "__main__":
    main()
