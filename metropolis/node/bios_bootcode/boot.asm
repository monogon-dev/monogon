org 7c00h

start:
	jmp main

; si: string data, null terminated
; di: start offset
writestring:
	mov al, [si]
	or al, al
	jz writestring_done
	inc si
	mov byte [fs:di], al
	add di, 2
	jmp writestring
writestring_done:
	ret

; si: rle encoded data (high bit == color, lower 7: length)
; di: start offset
writegfx:
	mov al, [si]
	or al, al
	jz writegfx_done
	inc si

	mov cl, al
	and cx, 0b01111111
	shr al, 7

writegfx_nextinner:
	or al, al
	jz writegfx_space
	mov byte [fs:di], 'M'
writegfx_space:
	add di, 2
	sub cx, 1
	jz writegfx
	jmp writegfx_nextinner
writegfx_done:
	ret

main:
	xor ax, ax
	mov ds, ax

	; set mode 3 (text 80x25, 16 color)
	mov ax, 0x3
	int 0x10

	; set up fs segment to point at framebuffer
	mov ax, 0xb800
	mov fs, ax

	mov di, 4
	mov si, logo
	call writegfx

	mov di, 3400
	mov si, line1
	call writestring

	mov di, 3544
	mov si, line2
	call writestring

end:
	jmp end

; Workaround to pass file as argument
%macro incdef 1
    %push _incdef_
	%defstr %$file %{1}
	%include %{$file}
	%pop
%endmacro

incdef LOGO

line1:
	db "Hi there! Didn't see you coming in.", 0

line2:
	db "Unfortunately, Metropolis can only boot in UEFI mode.", 0

db 0x55
db 0xAA

; We don't fill the rest with zeros, as this is done by mkimage and friends.