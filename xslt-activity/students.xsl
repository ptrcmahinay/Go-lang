<?xml version = "1.0" encoding = "UTF-8"?>
<xsl:stylesheet version = "1.0" xmlns:xsl = "http://www.w3.org/1999/XSL/Transform"> 

<xsl:template match="/">
  <html>
    <head>
      <link rel="stylesheet" type="text/css" href="students.css" />
    </head>
    <body>
      <h1>Students Data</h1>

      <table border="1">
        <tr>
          <th>Student ID</th>
          <th>Name</th>
          <th>Birthday</th>
          <th>Email</th>
          <th>Course</th>
          <th>Section</th>
        </tr>

        <xsl:for-each select="students/student"> 
          <tr>
            <td><xsl:value-of select="@id"/></td>
            <td><xsl:value-of select="name"/></td>
            <td><xsl:value-of select="birthday"/></td>
            <td><xsl:value-of select="email"/></td>
            <td><xsl:value-of select="department/course"/></td>
            <td><xsl:value-of select="department/section"/></td>
          </tr>
        </xsl:for-each>
      </table>
    </body>
  </html>
</xsl:template>  
</xsl:stylesheet>