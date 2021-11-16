package graphics

var vs2d = `#version 100

uniform mat3 mvp;
uniform mat2 uvp;

attribute  vec2 vert;
attribute vec2 clr_vert;
varying vec2 clr_TexCoords;


mat2 rotate2d(float _angle){
    return mat2(cos(_angle),-sin(_angle),
	                sin(_angle),cos(_angle));
}

mat2 scale(vec2 _scale){
    return mat2(_scale.x,0.0, 
				0.0,_scale.y);
					
}
mat4 projection(float vw, float vh){

	return  mat4(
	2.0 / vw, 0, 0, 0,
	0, 2.0 / vh, 0, 0,
	0, 0, 1, 0,
    -1, 1, 0, 1);

}


void main()
{
	
	
	float vw = mvp[0].x;
	float vh = mvp[1].x;
	float w = mvp[0].y; 
	float h = mvp[1].y; 
	float px = mvp[0].z;
	float py = mvp[1].z;
	float sx = mvp[2].x;
	float sy = mvp[2].y;

	float tw = uvp[0].x; 
	float th = uvp[0].y; 
	float minx = uvp[1].x; 
	float miny = uvp[1].y; 

	mat2 size = mat2(w, 0, 0, h);

	// Texture 
	//  +------------+
	//  | +-------+  |
	//  | |		  |  |
	//  | |		  |	 |
	//	| +-------+  |
	//  +------------+

	float upx = (minx)/tw;
	float upy = (miny)/tw;
	vec2 clr = clr_vert* scale(vec2(w/tw, h/th)) ;
	clr_TexCoords = vec2(clr.x + upx, clr.y + upy);

	// Rectangle 
	vec2 p = vert * scale(vec2(sx, sy)) * size;
	gl_Position =  projection(vw, vh) * vec4(p.x + px, p.y - py, 0, 1);
}

`

var fs2d = `#version 100

precision mediump float;

uniform sampler2D tex;
varying  vec2 clr_TexCoords;

void main()
{
	gl_FragColor = texture2D(tex, clr_TexCoords);
} 
`
