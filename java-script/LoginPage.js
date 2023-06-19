var tl= gsap.timLine({repeat: -1, repeatDelay: 1, yoyo: true})
tl
.from(splitText.chars, {duration: .5, opacity: 0, stagger: .125, ease: 'power1. In'})
.to(splitText.chars, {duration: .25, opacity: 0, stagger: .125, ease: 'power3. inOut'}, '+=3');